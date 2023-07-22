package v1

import (
	"context"
	"net/http"
	"strings"

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

	if err := dino.ParseSpecies(strings.Title(cdr.Species)); err != nil {
		e.Add("species", "is invalid")
	}

	if err := dino.ParseDiet(cdr.Diet); err != nil {
		e.Add("diet", "is invalid")
	}

	return e
}

// CreateDinoResponse - represents a client create dino response.
type CreateDinoResponse struct {
	Dinosaur ClientDino `json:"dinosaur"`
}

// CreateCage - invoked by POST /v1/dinosaurs.
func (c *Controller) CreateDino(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.Log.Info().Msg("Creating Dino.")

	var input CreateDinoRequest
	if err := api.Decode(r, &input); err != nil {
		c.Log.Err(err).Msg("Unable to decode create dino request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.Log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	d, err := c.Dino.Create(ctx, toCoreNewDino(input))
	if err != nil {
		c.Log.Err(err).Msg("Unable to create dino.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.Log.Info().Msg("Successfully created Dino.")
	return api.Respond(w, http.StatusCreated, CreateDinoResponse{Dinosaur: toClientDino(d)})
}

// ListDinoSpeciesResponse - represents list dino species response.
type ListDinoSpeciesResponse struct {
	DinoSpecies []ClientDinoSpecies `json:"species"`
}

// ListDinoSpecies - invoked by GET /v1/dinosaures/species.
func (c *Controller) ListDinoSpecies(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.Log.Info().Msg("Listing dino species.")

	out := make([]ClientDinoSpecies, 0, len(dino.DinoSpeciesMapping))
	for species, diet := range dino.DinoSpeciesMapping {
		out = append(out, ClientDinoSpecies{
			Species: species,
			Diet:    diet.String(),
		})
	}

	c.Log.Info().Msg("Successfully listed dino species.")
	return api.Respond(w, http.StatusOK, ListDinoSpeciesResponse{DinoSpecies: out})
}
