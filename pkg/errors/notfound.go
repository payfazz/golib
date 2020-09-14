package errors

import (
	"bytes"
	"fmt"
)

// NotFoundError , error containing validation message and properties.
type NotFoundError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (err NotFoundError) Error() string {
	b := new(bytes.Buffer)
	if err.Data == "" {
		fmt.Fprint(b, "not found")
	} else {
		fmt.Fprintf(b, `'%s' not found`, err.Data)
	}
	return b.String()
}

// NotFound ...
func NotFound(code, data string) error {
	err := NotFoundError{
		Code:    code,
		Message: fmt.Sprintf(`'%s' not found`, data),
		Data:    data,
	}
	return err
}
