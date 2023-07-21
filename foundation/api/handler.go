package api

import (
	"context"
	"net/http"
)

// Handler - represents an http handler with extra context and error handling.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
