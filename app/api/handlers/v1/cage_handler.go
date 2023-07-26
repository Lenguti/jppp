package v1

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/foundation/api"
)

// CreateCageRequest - represents input for creating a new cage.
type CreateCageRequest struct {
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}

func (ccr *CreateCageRequest) validate() *api.ValidationError {
	e := api.NewValidationError()
	if err := cage.ParseType(ccr.Type); err != nil {
		e.Add("type", "is invalid")
	}

	if err := cage.ParseStatus(ccr.Status); err != nil {
		e.Add("status", "is invalid")
	}

	return e
}

// CreateCageResponse - represents a client create cage response.
type CreateCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// CreateCage - invoked by POST /v1/cages.
func (c *Controller) CreateCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Cage.")

	var input CreateCageRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create cage request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	cge, err := c.Cage.Create(ctx, toCoreNewCage(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to create cage.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Cage.")
	return api.Respond(w, http.StatusCreated, CreateCageResponse{Cage: toClientCage(cge)})
}

// GetCageResponse - represents a client get cage response.
type GetCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// GetCage - invoked by GET /v1/cages/:id.
func (c *Controller) GetCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Fetching Cage.")

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	cge, err := c.Cage.Get(ctx, id)
	if err != nil {
		c.log.Err(err).Msg("Unable to fetch cage.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully fetched Cage.")
	return api.Respond(w, http.StatusOK, GetCageResponse{Cage: toClientCage(cge)})
}

// ListCagesResponse - represents a client list cages response.
type ListCagesResponse struct {
	Cages []ClientCage `json:"cages"`
}

// ListCages - invoked by GET /v1/cages.
func (c *Controller) ListCages(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Cages.")

	status := api.QueryParam(r, queryParamStatus)
	var filters []core.Filter
	if status != "" {
		if err := cage.ParseStatus(status); err != nil {
			c.log.Err(err).Msg("Invalid cage status filter.")
			return api.BadRequestError("Invalid cage status filter.", err, nil)
		}
		filters = append(filters, core.Filter{Key: queryParamStatus, Value: strings.ToUpper(status)})
	}

	cgs, err := c.Cage.List(ctx, filters...)
	if err != nil {
		c.log.Err(err).Msg("Unable to list cages.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully listed Cage.")
	return api.Respond(w, http.StatusOK, ListCagesResponse{Cages: toClientCages(cgs)})
}

// UpdateCageRequest - represents input for updating a cage.
type UpdateCageRequest struct {
	Status string `json:"status"`
}

func (ccr *UpdateCageRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if err := cage.ParseStatus(ccr.Status); err != nil {
		e.Add("status", "is invalid")
	}

	return e
}

// UpdateCageResponse - represents a client update cage response.
type UpdateCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// UpdateCage - invoked by PATCH /v1/cages/:id.
func (c *Controller) UpdateCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Updating Cage.")

	var input UpdateCageRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode update cage request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	cge, err := c.Cage.UpdateStatus(ctx, id, cage.Status(strings.ToUpper(input.Status)))
	if err != nil {
		c.log.Err(err).Msg("Unable to update cage.")
		switch {
		case errors.Is(err, core.ErrPowerDownCage):
			return api.BadRequestError(err.Error(), err, nil)
		case errors.Is(err, core.ErrNotFound):
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully updated Cage.")
	return api.Respond(w, http.StatusOK, UpdateCageResponse{Cage: toClientCage(cge)})
}

// AddDinosaurToCageResponse - represents a client add dino to cage response.
type AddDinosaurToCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// AddDinosaurToCage - invoked by PATCH /v1/cages/:id/dinosaurs/:id.
func (c *Controller) AddDinosaurToCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Adding Dinosaur to Cage.")

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid cage id.", err, nil)
	}

	dinoIdStr := api.PathParam(r, dinoIDPathParam)
	dinoID, err := uuid.Parse(dinoIdStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid dino id.")
		return api.BadRequestError("Invalid dinosaur id.", err, nil)
	}

	cge, err := c.Cage.AddDino(ctx, id, dinoID)
	if err != nil {
		c.log.Err(err).Msg("Unable to add dino to cage.")
		switch {
		case errors.Is(err, core.ErrInvalidCagePowerDown),
			errors.Is(err, core.ErrInvalidCageAtCapacity),
			errors.Is(err, core.ErrInvalidCageInvalidType),
			errors.Is(err, core.ErrInvalidCageInvalidSpecies):
			return api.BadRequestError(err.Error(), err, nil)
		case errors.Is(err, core.ErrNotFound):
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully added Dinosaur to Cage.")
	return api.Respond(w, http.StatusOK, AddDinosaurToCageResponse{Cage: toClientCage(cge)})
}

// RemoveDinosaurFromCageResponse - represents a client remove dino from cage response.
type RemoveDinosaurFromCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// RemoveDinosaurFromCage - invoked by DELETE /v1/cages/:id/dinosaurs/:id.
func (c *Controller) RemoveDinosaurFromCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Removing Dinosaur from Cage.")

	idStr := api.PathParam(r, idPathParam)
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid cage id.", err, nil)
	}

	dinoIdStr := api.PathParam(r, dinoIDPathParam)
	dinoID, err := uuid.Parse(dinoIdStr)
	if err != nil {
		c.log.Err(err).Msg("Invalid dino id.")
		return api.BadRequestError("Invalid dinosaur id.", err, nil)
	}

	cge, err := c.Cage.RemoveDino(ctx, id, dinoID)
	if err != nil {
		c.log.Err(err).Msg("Unable to remove dino from cage.")
		switch {
		case errors.Is(err, core.ErrInvalidCageInvalidRemoval):
			return api.BadRequestError(err.Error(), err, nil)
		case errors.Is(err, core.ErrNotFound):
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully removed Dinosaur from Cage.")
	return api.Respond(w, http.StatusOK, RemoveDinosaurFromCageResponse{Cage: toClientCage(cge)})
}
