package gocached

import (
	"context"
	"sync"
	"time"
)

type Options struct {
	DefaultTTL *time.Duration
}

func WithDefaultTTL(ttl time.Duration) func(*Options) {
	return func(options *Options) {
		options.DefaultTTL = &ttl
	}
}

func NewCache[K comparable, V any](
	evicter Evicter[K],
	optFns ...func(*Options),
) *Cache[K, V] {
	if evicter == nil {
		evicter = NewNoEvicter[K]()
	}

	options := &Options{}
	for _, optFn := range optFns {
		optFn(options)
	}

	return &Cache[K, V]{
		options: *options,
		data:    make(map[K]Entry[K, V]),
		lock:    sync.RWMutex{},
		evicter: evicter,
	}
}

type Cache[K comparable, V any] struct {
	options Options
	evicter Evicter[K]
	data    map[K]Entry[K, V]
	lock    sync.RWMutex
}

func (cache *Cache[K, V]) Set(ctx context.Context, key K, value V, opts ...EntryFn) error {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	return cache.set(ctx, key, value, opts...)
}

func (cache *Cache[K, V]) Get(ctx context.Context, key K) (V, bool, error) {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	return cache.get(ctx, key)
}

func (cache *Cache[K, V]) Del(ctx context.Context, key K) error {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	return cache.del(ctx, key)
}

func (cache *Cache[K, V]) Has(ctx context.Context, key K) bool {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	return cache.has(ctx, key)
}

func (cache *Cache[K, V]) set(ctx context.Context, key K, value V, opts ...EntryFn) error {
	entry := Entry[K, V]{
		Key:     key,
		Value:   value,
		Written: time.Now(),
		Options: makeEntryOptions(opts),
	}
	cache.data[key] = entry
	cache.evicter.Promote(ctx, key, 1)
	return nil
}

func (cache *Cache[K, V]) getEntry(ctx context.Context, key K) (Entry[K, V], bool) {
	if cache.has(ctx, key) {
		entry := cache.data[key]
		return entry, true
	}
	return Entry[K, V]{}, false
}

func (cache *Cache[K, V]) get(ctx context.Context, key K) (V, bool, error) {
	entry, ok := cache.getEntry(ctx, key)

	if !ok {
		return *new(V), false, nil
	}

	entryIsExplicitlyOutdated := entry.IsOutdated()
	entryIsExpiredDueToDefaultTTL := cache.options.DefaultTTL != nil && time.Now().Sub(entry.Written) > *cache.options.DefaultTTL

	if entryIsExplicitlyOutdated || entryIsExpiredDueToDefaultTTL {
		cache.del(ctx, key)
		return *new(V), false, nil
	}

	return entry.Value, true, nil
}

func (cache *Cache[K, V]) del(ctx context.Context, key K) error {
	delete(cache.data, key)
	cache.evicter.Evict(ctx, key)
	return nil
}

func (cache *Cache[K, V]) has(ctx context.Context, key K) bool {
	_, ok := cache.data[key]
	cache.evicter.Promote(ctx, key, 1)
	return ok
}
