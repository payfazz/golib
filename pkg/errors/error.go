package errors

import (
	"bytes"
	"fmt"
)

// Error , error containing validation message and properties.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err Error) Error() string {
	b := new(bytes.Buffer)
	code := err.Code
	if code == "" {
		code = "ERROR"
	}
	message := err.Message
	if message == "" {
		message = "error has occurred"
	}
	fmt.Fprintf(b, `%s - %s.`, code, message)
	return b.String()
}

// New ...
func New(code, message string) error {
	err := NotFoundError{
		Code:    code,
		Message: message,
	}
	return err
}
