package v1

import (
	"context"
	"fmt"
	"net/http"
)

// Routes - route definitions for v1.
func Routes(c *Controller) {
	const version = "v1"

	c.Router.Handle(http.MethodGet, version, "/status", c.status)

	c.Router.Handle(http.MethodPost, version, "/cages", c.CreateCage)
	c.Router.Handle(http.MethodGet, version, "/cages", c.ListCages)
	c.Router.Handle(http.MethodGet, version, "/cages/:id", c.GetCage)

	c.Router.Handle(http.MethodGet, version, "/dinosaurs/species", c.ListDinoSpecies)
	c.Router.Handle(http.MethodPost, version, "/dinosaurs", c.CreateDino)
}

func (c *Controller) status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := c.db.Connect(); err != nil {
		return fmt.Errorf("status: unable to connect to db: %w", err)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "ok"}`))
	return nil
}
