package gocached_test

import (
	"context"
	"testing"
	"time"

	"github.com/lazharichir/gocached"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	// Initialize cache
	c := gocached.NewCache[string, string](nil)

	// Test Set and Get
	c.Set(ctx, "key1", "value1")
	c.Set(ctx, "key2", "value2")

	v1, found, err := c.Get(ctx, "key1")
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "value1", v1)

	v2, found, err := c.Get(ctx, "key2")
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "value2", v2)

	// Test Get with non-existing key
	v3, found, err := c.Get(ctx, "key3")
	assert.NoError(t, err)
	assert.False(t, found)
	assert.Equal(t, *new(string), v3)

	// Test Has
	assert.True(t, c.Has(ctx, "key1"))
	assert.False(t, c.Has(ctx, "key3"))

	// Test Del
	c.Del(ctx, "key1")
	assert.False(t, c.Has(ctx, "key1"))
}

func TestDefaultTTL(t *testing.T) {
	ctx := context.Background()
	ttl := 250 * time.Millisecond
	cache := gocached.NewCache[string, int](nil, gocached.WithDefaultTTL(ttl))
	cache.Set(ctx, "1", 1)

	v1, ok, _ := cache.Get(ctx, "1")
	assert.True(t, ok, "key 1 should be present")
	assert.Equal(t, 1, v1, "key 1 should be present")

	time.Sleep(ttl + (20 * time.Millisecond))

	v1, ok, _ = cache.Get(ctx, "1")
	assert.False(t, ok, "key 1 should have expired")
	assert.Equal(t, 0, v1, "key 1 should have expired")
}

func TestEntryWithTTL(t *testing.T) {
	ttl := 100 * time.Millisecond
	ctx := context.Background()
	cache := gocached.NewCache[string, int](nil)
	cache.Set(ctx, "1", 1)
	cache.Set(ctx, "2", 2, gocached.WithTTL(ttl))

	v1, ok, _ := cache.Get(ctx, "1")
	assert.True(t, ok, "key 1 should be present")
	assert.Equal(t, 1, v1, "key 1 should be present")

	v2, ok, _ := cache.Get(ctx, "2")
	assert.True(t, ok, "key 2 should be present")
	assert.Equal(t, 2, v2, "key 2 should be present")

	time.Sleep(ttl + (20 * time.Millisecond))

	v1, ok, _ = cache.Get(ctx, "1")
	assert.True(t, ok, "key 1 should still be present")
	assert.Equal(t, 1, v1, "key 1 should still be present")

	v2, ok, _ = cache.Get(ctx, "2")
	assert.False(t, ok, "key 2 should have expired")
	assert.Equal(t, 0, v2, "key 2 should have expired")
}

func TestEntryWithExpiryDate(t *testing.T) {
	ctx := context.Background()
	expiryDate := time.Now().Add(100 * time.Millisecond)
	cache := gocached.NewCache[string, int](nil)
	cache.Set(ctx, "1", 1)
	cache.Set(ctx, "2", 2, gocached.WithExpiryDate(expiryDate))

	v1, ok, _ := cache.Get(ctx, "1")
	assert.True(t, ok, "key 1 should be present")
	assert.Equal(t, 1, v1, "key 1 should be present")

	v2, ok, _ := cache.Get(ctx, "2")
	assert.True(t, ok, "key 2 should be present")
	assert.Equal(t, 2, v2, "key 2 should be present")

	time.Sleep(100 * time.Millisecond)

	v1, ok, _ = cache.Get(ctx, "1")
	assert.True(t, ok, "key 1 should still be present")
	assert.Equal(t, 1, v1, "key 1 should still be present")

	v2, ok, _ = cache.Get(ctx, "2")
	assert.False(t, ok, "key 2 should have expired")
	assert.Equal(t, 0, v2, "key 2 should have expired")
}
