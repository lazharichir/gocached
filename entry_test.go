package gocached_test

import (
	"testing"
	"time"

	"github.com/lazharichir/gocached"
	"github.com/stretchr/testify/assert"
)

func TestEntry(t *testing.T) {
	expiryDate := time.Now().Add(1 * time.Hour)
	ttl := 10 * time.Minute
	entry := &gocached.Entry[string, int]{
		Key:     "foo",
		Value:   42,
		Written: time.Now(),
		Options: gocached.EntryOptions{
			ExpiryDate: &expiryDate,
			TTL:        &ttl,
		},
	}

	// Test IsOutdated method
	assert.False(t, entry.IsOutdated(), "Expected entry to not be outdated")

	// Test isPastExpiryDate method
	assert.False(t, entry.IsPastExpiryDate(), "Expected entry to not be past expiry date")

	// Test isPastTTL method
	assert.False(t, entry.IsPastTTL(), "Expected entry to not be past TTL")
}

func TestEntryOptions(t *testing.T) {
	expiryDate := time.Now().Add(1 * time.Hour)
	ttl := 10 * time.Minute

	// Test EntryOptions
	options := gocached.EntryOptions{
		ExpiryDate: &expiryDate,
		TTL:        &ttl,
	}

	assert.Equal(t, expiryDate, *options.ExpiryDate, "Expected expiry date to be %v, got %v", expiryDate, *options.ExpiryDate)
	assert.Equal(t, ttl, *options.TTL, "Expected TTL to be %v, got %v", ttl, *options.TTL)
}
