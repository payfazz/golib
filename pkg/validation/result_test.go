package validation_test

import (
	"testing"

	"github.com/payfazz/golib/pkg/validation"
)

func TestValidationResultAddSameField(t *testing.T) {
	name := "field1"
	r := validation.NewResult()
	r.Add(name, "error1")
	r.Add(name, "error2")
	if !r.HasErrors() {
		t.Error("expected value:'true' as result of ValidationResult.HasErrors()")
	}
	message := r.Get(name)
	if message != "error2" {
		t.Error("expected value:'error2' as result of ValidationResult.Get(name)")
	}
}
