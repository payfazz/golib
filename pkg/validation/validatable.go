package validation

import "context"

// Validatable ...
type Validatable interface {
	Validate(context context.Context) *Result
}
