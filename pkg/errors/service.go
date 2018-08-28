package errors

import (
	"bytes"
	"fmt"
)

// Service , error indicating service error.
type Service struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Error ...
func (err Service) Error() string {
	b := new(bytes.Buffer)
	if err.Name == "" {
		fmt.Fprintf(b, "service error")
	} else {
		fmt.Fprintf(b, `error on service '%s' : %s`, err.Name, err.Message)
	}
	return b.String()
}

// ServiceError ...
func ServiceError(name string, err error) error {
	errService := Service{
		Name:    name,
		Message: err.Error(),
	}
	return errService
}
