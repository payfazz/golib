package actor_test

import (
	"context"
	"testing"

	"github.com/payfazz/golib/pkg/actor"
)

func TestActorContext(t *testing.T) {
	name := "golib/actor"
	ctx := actor.NewContext(context.TODO(), name)

	act, ok := actor.FromContext(ctx)
	if !ok {
		t.Error("failed to get actor from context")
	}
	if act != name {
		t.Error("actor is invalid")
	}
}
