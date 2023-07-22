package v1

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/core/dino"
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
	c.Log.Info().Msg("Creating Cage.")

	var input CreateCageRequest
	if err := api.Decode(r, &input); err != nil {
		c.Log.Err(err).Msg("Unable to decode create cage request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.Log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	cge, err := c.Cage.Create(ctx, toCoreNewCage(input))
	if err != nil {
		c.Log.Err(err).Msg("Unable to create cage.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.Log.Info().Msg("Successfully created Cage.")
	return api.Respond(w, http.StatusCreated, CreateCageResponse{Cage: toClientCage(cge, nil)})
}

// GetCageResponse - represents a client get cage response.
type GetCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// GetCage - invoked by GET /v1/cages/:id.
func (c *Controller) GetCage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.Log.Info().Msg("Fetching Cage.")

	idStr := api.PathParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Log.Err(err).Msg("Invalid cage id.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	cge, err := c.Cage.Get(ctx, id)
	if err != nil {
		c.Log.Err(err).Msg("Unable to fetch cage.")
		return api.InternalServerError("Error.", err, nil)
	}

	var dinos []dino.Dinosaur
	if cge.CurrentCapacity > 0 {
		dinos, err = c.Dino.ListByCageID(ctx, cge.ID)
		if err != nil {
			c.Log.Err(err).Msg("Unable to list dinos.")
			return api.InternalServerError("Error.", err, nil)
		}
	}

	c.Log.Info().Msg("Successfully fetched Cage.")
	return api.Respond(w, http.StatusOK, GetCageResponse{Cage: toClientCage(cge, dinos)})
}

// ListCagesResponse - represents a client list cages response.
type ListCagesResponse struct {
	Cages []ClientCage `json:"cages"`
}

// ListCages - invoked by GET /v1/cages.
func (c *Controller) ListCages(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.Log.Info().Msg("Listing Cages.")

	cgs, err := c.Cage.List(ctx)
	if err != nil {
		c.Log.Err(err).Msg("Unable to list cages.")
		return api.InternalServerError("Error.", err, nil)
	}

	cageIDs := make([]uuid.UUID, 0, len(cgs))
	for _, v := range cgs {
		cageIDs = append(cageIDs, v.ID)
	}
	dinos, err := c.Dino.ListByCageIDs(ctx, cageIDs...)
	if err != nil {
		c.Log.Err(err).Msg("Unable to list dinos.")
		return api.InternalServerError("Error.", err, nil)
	}

	out := toClientCages(cgs, dinos)
	c.Log.Info().Msg("Successfully listed Cage.")
	return api.Respond(w, http.StatusOK, ListCagesResponse{Cages: out})
}
