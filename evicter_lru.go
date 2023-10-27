package gocached

import (
	"context"
	"sort"
	"time"
)

type LRUEvicter[K comparable] struct {
	internal map[K]time.Time
}

func NewLRUEvicter[K comparable]() *LRUEvicter[K] {
	return &LRUEvicter[K]{
		internal: make(map[K]time.Time),
	}
}

func (evicter *LRUEvicter[K]) Promote(ctx context.Context, key K, value int) {
	evicter.internal[key] = time.Now()
}

func (evicter *LRUEvicter[K]) Demote(ctx context.Context, key K, value int) {
	evicter.Demote(ctx, key, value)
}

func (evicter *LRUEvicter[K]) Evict(ctx context.Context, key K) {
	delete(evicter.internal, key)
}

func (evicter *LRUEvicter[K]) Evictees(ctx context.Context, n int) []K {
	// Create a slice to store the keys
	keys := make([]K, 0, len(evicter.internal))
	for key := range evicter.internal {
		keys = append(keys, key)
	}

	// Sort the keys by their internal
	sort.Slice(keys, func(i, j int) bool {
		return evicter.internal[keys[i]].Before(evicter.internal[keys[j]])
	})

	// Return the first n keys
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}
