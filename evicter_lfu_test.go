package gocached

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUEvicter(t *testing.T) {
	ctx := context.Background()
	lfu := NewLFUEvicter[string]()

	lfu.Promote(ctx, "key3", 1)
	lfu.Promote(ctx, "key3", 1)
	lfu.Promote(ctx, "key3", 1)

	lfu.Promote(ctx, "key2", 1)
	lfu.Promote(ctx, "key2", 1)

	lfu.Promote(ctx, "key1", 1)

	assert.Equal(t, []string{"key1"}, lfu.Evictees(ctx, 1))

	lfu.Promote(ctx, "key1", 5)

	assert.Equal(t, []string{"key2"}, lfu.Evictees(ctx, 1))

	lfu.Demote(ctx, "key1", 1)
	lfu.Demote(ctx, "key2", 1)
	lfu.Demote(ctx, "key3", 1)

	assert.Equal(t, []string{"key2"}, lfu.Evictees(ctx, 1))

	lfu.Demote(ctx, "key3", 3)
	lfu.Demote(ctx, "key2", 1)

	assert.Equal(t, []string{"key3"}, lfu.Evictees(ctx, 1))

	assert.ElementsMatch(t, []string{"key3", "key2"}, lfu.Evictees(ctx, 2))

	lfu.Evict(ctx, "key3")

	assert.ElementsMatch(t, []string{"key2", "key1"}, lfu.Evictees(ctx, 2))

	lfu.Evict(ctx, "key1")
	lfu.Evict(ctx, "key2")

	assert.ElementsMatch(t, []string{}, lfu.Evictees(ctx, 2))
}
