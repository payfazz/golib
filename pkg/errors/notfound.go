package errors

import (
	"bytes"
	"fmt"
)

// NotFoundError , error containing validation message and properties.
type NotFoundError struct {
	Data    string `json:"data"`
	Message string `json:"message"`
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
func NotFound(data string) error {
	err := NotFoundError{
		Data:    data,
		Message: fmt.Sprintf(`'%s' not found`, data),
	}
	return err
}
