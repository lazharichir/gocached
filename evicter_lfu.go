package gocached

import (
	"context"
	"sort"
)

type LFUEvicter[K comparable] struct {
	frequency map[K]int
}

func NewLFUEvicter[K comparable]() *LFUEvicter[K] {
	return &LFUEvicter[K]{
		frequency: make(map[K]int),
	}
}

func (evicter *LFUEvicter[K]) Promote(ctx context.Context, key K, value int) {
	evicter.frequency[key] = evicter.frequency[key] + value
}

func (evicter *LFUEvicter[K]) Demote(ctx context.Context, key K, value int) {
	evicter.frequency[key] = evicter.frequency[key] - value
}

func (evicter *LFUEvicter[K]) Evict(ctx context.Context, key K) {
	delete(evicter.frequency, key)
}

func (evicter *LFUEvicter[K]) Evictees(ctx context.Context, n int) []K {
	// Create a slice to store the keys
	keys := make([]K, 0, len(evicter.frequency))
	for key := range evicter.frequency {
		keys = append(keys, key)
	}

	// Sort the keys by their frequency
	sort.Slice(keys, func(i, j int) bool {
		return evicter.frequency[keys[i]] < evicter.frequency[keys[j]]
	})

	// Return the first n keys
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}
