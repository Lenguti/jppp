package api

import (
	"encoding/json"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// Decode - decodes request body into the provided value.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(val)
}

// PathParam - returns the path parameters from the request.
func PathParam(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}
