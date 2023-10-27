package gocached_test

import (
	"context"
	"testing"

	"github.com/lazharichir/gocached"
	"github.com/stretchr/testify/assert"
)

func TestNoEvicter(t *testing.T) {
	ctx := context.Background()
	ev := gocached.NewNoEvicter[any]()
	assert.NotNil(t, ev)
	assert.NotPanics(t, func() {
		ev.Promote(ctx, 1, 0)
		ev.Demote(ctx, 1, 0)
		ev.Evictees(ctx, 1)
		ev.Evict(ctx, 1)
	})
}
