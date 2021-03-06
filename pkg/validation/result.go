package validation

import (
	"fmt"
)

// Result ...
type Result struct {
	Fields []*Field `json:"fields"`
}

// Field ...
type Field struct {
	Name    string `json:"field"`
	Message string `json:"message"`
}

// NewResult return new validation result
func NewResult() *Result {
	result := &Result{}
	return result
}

// Add , append new Field if name does not exist, updating existing message if exists
func (r *Result) Add(name string, message string) {
	f := r.findField(name)
	if f == nil {
		f = &Field{Name: name}
		r.Fields = append(r.Fields, f)
	}
	f.Message = message
}

// Get , return error message
func (r *Result) Get(name string) string {
	f := r.findField(name)
	if f == nil {
		return ""
	}
	return f.Message
}

// AddInvalid , append result with predefined invalid message
func (r *Result) AddInvalid(name string) {
	r.Add(name, fmt.Sprintf(`value for '%s' is invalid`, name))
}

// AddRequired , append result with predefined invalid message
func (r *Result) AddRequired(name string) {
	r.Add(name, fmt.Sprintf(`value for '%s' is required`, name))
}

// HasErrors , checks if result has error
func (r *Result) HasErrors() bool {
	return len(r.Fields) > 0
}

// Error , return error object
func (r *Result) Error(code string) error {
	return NewError(code, r)
}

func (r *Result) findField(name string) *Field {
	for _, f := range r.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}
