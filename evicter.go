package gocached

import (
	"context"
)

type Evicter[K comparable] interface {
	// Promote promotes the key.
	Promote(ctx context.Context, key K, value int)
	// Demote demotes the key.
	Demote(ctx context.Context, key K, value int)
	// Evictees returns the keys to evict.
	Evictees(ctx context.Context, n int) []K
	// Evict removes the key from the evicter and forgets about it completely.
	Evict(ctx context.Context, key K)
}
