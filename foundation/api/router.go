package api

import (
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/rs/zerolog"
)

// Router - represents api router.
type Router struct {
	Log zerolog.Logger

	mux *httptreemux.ContextMux
}

// NewRouter - initialized new router.
func NewRouter(log zerolog.Logger) *Router {
	m := httptreemux.NewContextMux()

	m.GET("/healthcheck", healthCheck)

	return &Router{
		Log: log,

		mux: m,
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// ServeHTTP - satisfies the http.Handler interface.
func (rr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.mux.ServeHTTP(w, r)
}

// Handle - a wrapper around default http.Handle.
func (rr *Router) Handle(method, group, path string, h Handler) {
	p := fmt.Sprintf("/%s/%s", group, path)
	rr.handle(method, p, h)
}

func (rr *Router) handle(method string, path string, h Handler) {
	hh := func(w http.ResponseWriter, r *http.Request) {
		if err := h(r.Context(), w, r); err != nil {
			rr.Log.Error().Err(err).Msg("Error handling request.")
		}
	}
	rr.mux.Handle(method, path, hh)
}
