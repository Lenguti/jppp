package handlers

import (
	"net/http"

	v1 "github.com/lenguti/jppp/app/api/handlers/v1"
	"github.com/lenguti/jppp/foundation/api"
	"github.com/rs/zerolog"
)

// NewV1Handler - initializes new v1 handler with provided config.
func NewV1Handler(log zerolog.Logger, cfg v1.Config) http.Handler {
	r := api.NewRouter(log)
	v1.Routes(r, cfg)
	return r
}
