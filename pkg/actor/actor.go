package actor

import (
	"context"
)

type key int

const actorKey key = 0

// NewContext , return new context with actor information.
func NewContext(ctx context.Context, actor string) context.Context {
	return context.WithValue(ctx, actorKey, actor)
}

// FromContext , return actor from a context.
func FromContext(ctx context.Context) (string, bool) {
	actor, ok := ctx.Value(actorKey).(string)
	return actor, ok
}
