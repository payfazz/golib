package validation

import (
	"bytes"
	"fmt"
)

// Error , error containing validation message and properties.
type Error struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Fields  []*Field `json:"fields"`
}

func (err Error) Error() string {
	b := new(bytes.Buffer)
	if err.Message == "" {
		fmt.Fprint(b, "validation error")
	} else {
		fmt.Fprint(b, err.Message)
	}
	if err.Fields != nil && len(err.Fields) > 0 {
		fmt.Fprint(b, "\nPlease check following fields:\n")
		for _, field := range err.Fields {
			fmt.Fprintf(b, "%s : %s\n", field.Name, field.Message)
		}
	}
	return b.String()
}

// NewError , return new validation error instance.
func NewError(code string, result *Result) error {
	err := Error{
		Code:    code,
		Message: "validation error",
		Fields:  result.Fields,
	}
	return err
}
