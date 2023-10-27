package gocached

import (
	"context"
)

type Cacher[K comparable, V any] interface {
	Set(ctx context.Context, key K, value V, opts ...EntryFn) error
	Get(ctx context.Context, key K) (V, bool, error)
	Has(ctx context.Context, key K) bool
	Del(ctx context.Context, key K) error
}
