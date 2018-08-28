package validation

// Validatable ...
type Validatable interface {
	Validate(context interface{}) *Result
}
