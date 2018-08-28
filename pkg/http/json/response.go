package json

import (
	"encoding/json"
	"net/http"

	e "github.com/payfazz/golib/pkg/errors"
	"github.com/payfazz/golib/pkg/validation"
)

func encodeBody(writer http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(writer).Encode(data)
}

// ResponseWithData , response with data
func ResponseWithData(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)
	if data != nil {
		encodeBody(writer, data)
	}
}

// ResponseWithError , response with error
func ResponseWithError(writer http.ResponseWriter, err error) {
	switch err.(type) {
	case validation.Error:
		ResponseWithData(writer, http.StatusUnprocessableEntity, err)
	case e.NotFoundError:
		ResponseWithData(writer, http.StatusNotFound, err)
	default:
		httpCode := http.StatusBadRequest
		http.Error(writer, err.Error(), httpCode)
	}
}
