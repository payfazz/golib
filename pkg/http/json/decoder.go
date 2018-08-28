package json

import (
	"encoding/json"
	"net/http"
)

// Decode ...
func Decode(request *http.Request, i interface{}) error {
	defer request.Body.Close()
	return json.NewDecoder(request.Body).Decode(i)
}
