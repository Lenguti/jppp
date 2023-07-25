package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/lenguti/jppp/foundation/api"
)

// CreateDinoRequest - represents input for creating a new dinosaur.
type CreateDinoRequest struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Diet    string `json:"diet"`
}

func (cdr *CreateDinoRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if cdr.Name == "" {
		e.Add("name", "is required")
	}

	dinoDiet, err := dino.ParseSpecies(cdr.Species)
	if err != nil {
		e.Add("species", "is invalid")
	}

	if err := dino.ParseDiet(cdr.Diet); err != nil {
		e.Add("diet", "is invalid")
	}

	if dinoDiet != dino.Diet(cdr.Diet) {
		e.Add("species diet", "is invalid")
	}

	return e
}

// CreateDinoResponse - represents a client create dino response.
type CreateDinoResponse struct {
	Dinosaur ClientDino `json:"dinosaur"`
}

// CreateCage - invoked by POST /v1/dinosaurs.
func (c *Controller) CreateDino(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Dino.")

	var input CreateDinoRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create dino request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	d, err := c.Dino.Create(ctx, toCoreNewDino(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to create dino.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Dino.")
	return api.Respond(w, http.StatusCreated, CreateDinoResponse{Dinosaur: toClientDino(d)})
}

// ListDinoSpeciesResponse - represents list dino species response.
type ListDinoSpeciesResponse struct {
	DinoSpecies []ClientDinoSpecies `json:"species"`
}

// ListDinoSpecies - invoked by GET /v1/dinosaures/species.
func (c *Controller) ListDinoSpecies(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing dino species.")

	out := make([]ClientDinoSpecies, 0, len(dino.DinoSpeciesMapping))
	for species, diet := range dino.DinoSpeciesMapping {
		out = append(out, ClientDinoSpecies{
			Species: species,
			Diet:    diet.String(),
		})
	}

	c.log.Info().Msg("Successfully listed dino species.")
	return api.Respond(w, http.StatusOK, ListDinoSpeciesResponse{DinoSpecies: out})
}

// GetDinoResponse - represents a client get dino response.
type GetDinoResponse struct {
	Dinosaur ClientDino `json:"dinosaur"`
}

// GetDino - invoked by GET /v1/dinosaurs/:id.
func (c *Controller) GetDino(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Fetching Dino.")

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid dino id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	d, err := c.Dino.Get(ctx, id)
	if err != nil {
		c.log.Err(err).Msg("Unable to fetch dino.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully fetched Dino.")
	return api.Respond(w, http.StatusOK, GetDinoResponse{Dinosaur: toClientDino(d)})
}

// ListDinosResponse - represents a client list dinos response.
type ListDinosResponse struct {
	Dinosaurs []ClientDino `json:"dinosaurs"`
}

// ListDinos - invoked by GET /v1/dinosaurs.
func (c *Controller) ListDinos(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Dinos.")

	ds, err := c.Dino.List(ctx)
	if err != nil {
		c.log.Err(err).Msg("Unable to list dinos.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully listed Dinos.")
	return api.Respond(w, http.StatusOK, ListDinosResponse{Dinosaurs: toClientDinos(ds)})
}

// UpdateDinoRequest - represents input for updating a dinosaur.
type UpdateDinoRequest struct {
	Name string `json:"name"`
}

func (udr *UpdateDinoRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if udr.Name == "" {
		e.Add("name", "is required")
	}

	return e
}

// UpdateDinoResponse - represents a client update dino response.
type UpdateDinoResponse struct {
	Dinosaur ClientDino `json:"dinosaur"`
}

// UpdateDino - invoked by PATCH /v1/dinosaurs/:id.
func (c *Controller) UpdateDino(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Updating Dinosaur.")

	var input UpdateDinoRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode update dino request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid dino id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	d, err := c.Dino.UpdateName(ctx, id, input.Name)
	if err != nil {
		c.log.Err(err).Msg("Unable to update dino.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully updated Dinosaur.")
	return api.Respond(w, http.StatusOK, UpdateDinoResponse{Dinosaur: toClientDino(d)})
}

// ListCageDinosaursResponse - represents a client list cage dinosaurs response.
type ListCageDinosaursResponse struct {
	Dinosaurs []ClientDino `json:"dinosaurs"`
}

// ListCageDinosaurs - invoked by GET /v1/cages/:id/dinosaurs.
func (c *Controller) ListCageDinosaurs(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Dinosaurs for Cage.")

	cageIDStr := api.PathParam(r, idPathParam)
	cageID, err := uuid.Parse(cageIDStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid cage id.", err, nil)
	}

	species := api.QueryParam(r, queryParamSpecies)
	var filters []core.Filter
	if species != "" {
		filters = append(filters, core.Filter{Key: queryParamSpecies, Value: strings.Title(species)})
	}

	dns, err := c.Dino.ListByCageID(ctx, cageID, filters...)
	if err != nil {
		c.log.Err(err).Msg("Unable to list dinosaurs for Cage.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully listed Dinosaurs for Cage.")
	return api.Respond(w, http.StatusOK, ListCageDinosaursResponse{Dinosaurs: toClientDinos(dns)})
}
