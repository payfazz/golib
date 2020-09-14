package validation_test

import (
	"testing"

	"github.com/payfazz/golib/pkg/validation"
)

func TestWithEmptyMessageAndProperties(t *testing.T) {
	r := validation.NewResult()
	instance := r.Error("C23")
	errMessage := instance.Error()
	if errMessage != "validation error" {
		t.Fatalf("error message should be 'validation error' when creating ValidationError with empty message and properties")
	}
}
