package validation

import (
	"bytes"
	"fmt"
)

// Error , error containing validation message and properties.
type Error struct {
	Result *Result
}

func (err Error) Error() string {
	b := new(bytes.Buffer)
	if err.Result.Message == "" {
		fmt.Fprint(b, "validation error")
	} else {
		fmt.Fprint(b, err.Result.Message)
	}
	if err.Result.Fields != nil && len(err.Result.Fields) > 0 {
		fmt.Fprint(b, "\nPlease check following fields:\n")
		for _, field := range err.Result.Fields {
			fmt.Fprintf(b, "%s : %s\n", field.Name, field.Message)
		}
	}
	return b.String()
}

// NewError , return new validation error instance.
func NewError(r *Result) error {
	valErr := Error{Result: r}
	return valErr
}
