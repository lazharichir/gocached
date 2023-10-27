package gocached

import "context"

type NoEvicter[K comparable] struct{}

func NewNoEvicter[K comparable]() *NoEvicter[K] {
	return &NoEvicter[K]{}
}

func (evicter *NoEvicter[K]) Promote(ctx context.Context, key K, delta int) {}

func (evicter *NoEvicter[K]) Demote(ctx context.Context, key K, delta int) {}

func (evicter *NoEvicter[K]) Evict(ctx context.Context, key K) {}

func (evicter *NoEvicter[K]) Evictees(ctx context.Context, n int) []K {
	return []K{}
}
