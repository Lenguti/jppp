package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/lenguti/jppp/foundation/api"
)

// Routes - route definitions for v1.
func Routes(router *api.Router, cfg Config) {
	const version = "v1"
	c := controller{
		cfg: cfg,
		log: router.Log,
	}

	router.Handle(http.MethodGet, version, "status", c.status)
}

func (c controller) status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.cfg.DBUser, c.cfg.DBPass),
		Host:     "db:5432",
		Path:     c.cfg.DBName,
		RawQuery: q.Encode(),
	}

	if _, err := sqlx.Connect("postgres", u.String()); err != nil {
		return fmt.Errorf("db connect: unable to connect to db: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "ok"}`))
	return nil
}
