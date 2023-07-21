package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
