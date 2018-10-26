package errors

import (
	"bytes"
	"fmt"
)

// ServiceError , error indicating service error.
type ServiceError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

// Error ...
func (err ServiceError) Error() string {
	b := new(bytes.Buffer)
	if err.Name == "" {
		fmt.Fprintf(b, "service error")
	} else {
		fmt.Fprintf(b, `error on service '%s' : %s`, err.Name, err.Message)
	}
	return b.String()
}

// Service ...
func Service(code, name string, err error) error {
	errService := ServiceError{
		Code:    code,
		Message: err.Error(),
		Name:    name,
	}
	return errService
}
