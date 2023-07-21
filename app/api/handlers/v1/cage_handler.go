package v1

import (
	"context"
	"net/http"

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

// CreateCageResponse - represents a client cage response.
type CreateCageResponse struct {
	Cage ClientCage `json:"cage"`
}

// CreateCage - invoked by /v1/cages.
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
	return api.Respond(w, http.StatusCreated, toClientCage(cge))
}
