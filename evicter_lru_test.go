package gocached

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUEvicter(t *testing.T) {
	ctx := context.Background()
	lfu := NewLRUEvicter[string]()

	lfu.Promote(ctx, "key1", 0)
	lfu.Promote(ctx, "key2", 0)
	lfu.Promote(ctx, "key3", 0)

	assert.Equal(t, []string{"key1"}, lfu.Evictees(ctx, 1))

	lfu.Promote(ctx, "key1", 0)

	assert.Equal(t, []string{"key2"}, lfu.Evictees(ctx, 1))
	assert.Equal(t, []string{"key2", "key3"}, lfu.Evictees(ctx, 2))

}
