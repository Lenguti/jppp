package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// Decode - decodes request body into the provided value.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(val)
}

// Respond - creates a response with desired statusCode and payload.
func Respond(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.WriteHeader(statusCode)
	if v != nil {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			return fmt.Errorf("respond: failed to encode response: %w", err)
		}
	}
	return nil
}

// PathParam - returns the path parameters from the request.
func PathParam(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

// QueryParam returns the query parameters from the request.
func QueryParam(r *http.Request, key string) string {
	q := r.URL.Query()
	return q.Get(key)
}
