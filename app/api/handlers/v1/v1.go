package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lenguti/jppp/foundation/api"
)

const (
	idPathParam     = "id"
	dinoIDPathParam = "dinoId"
)

const (
	queryParamStatus  = "status"
	queryParamSpecies = "species"
)

// Routes - route definitions for v1.
func (c *Controller) Routes() *api.Router {
	const version = "v1"

	c.router.Handle(http.MethodGet, version, "/status", c.status)

	c.router.Handle(http.MethodPost, version, "/cages", c.CreateCage)
	c.router.Handle(http.MethodGet, version, "/cages", c.ListCages)
	c.router.Handle(http.MethodGet, version, "/cages/:id", c.GetCage)
	c.router.Handle(http.MethodPatch, version, "/cages/:id", c.UpdateCage)
	c.router.Handle(http.MethodPatch, version, "/cages/:id/dinosaurs/:dinoId", c.AddDinosaurToCage)
	c.router.Handle(http.MethodDelete, version, "/cages/:id/dinosaurs/:dinoId", c.RemoveDinosaurFromCage)
	c.router.Handle(http.MethodGet, version, "/cages/:id/dinosaurs", c.ListCageDinosaurs)

	c.router.Handle(http.MethodGet, version, "/dinosaurs/species", c.ListDinoSpecies)
	c.router.Handle(http.MethodPost, version, "/dinosaurs", c.CreateDino)
	c.router.Handle(http.MethodGet, version, "/dinosaurs", c.ListDinos)
	c.router.Handle(http.MethodGet, version, "/dinosaurs/:id", c.GetDino)
	c.router.Handle(http.MethodPatch, version, "/dinosaurs/:id", c.UpdateDino)

	return c.router
}

func (c *Controller) status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := c.db.Connect(); err != nil {
		return fmt.Errorf("status: unable to connect to db: %w", err)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "ok"}`))
	return nil
}
