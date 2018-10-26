package errors

import (
	"bytes"
	"fmt"
)

// PermissionError , error regarding permission.
type PermissionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err PermissionError) Error() string {
	b := new(bytes.Buffer)
	if err.Message == "" {
		fmt.Fprint(b, "permission error")
	} else {
		fmt.Fprintf(b, `permission error: '%s'`, err.Message)
	}
	return b.String()
}

// Permission ...
func Permission(code, message string) error {
	err := PermissionError{
		Code:    code,
		Message: message,
	}
	return err
}
