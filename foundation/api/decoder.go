package api

import (
	"encoding/json"
	"net/http"
)

// Decode - decodes request body into the provided value.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(val)
}
