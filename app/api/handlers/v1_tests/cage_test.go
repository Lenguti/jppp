package v1_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dimfeld/httptreemux"
	"github.com/google/uuid"
	v1 "github.com/lenguti/jppp/app/api/handlers/v1"
	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/lenguti/jppp/foundation/api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCage(t *testing.T) {
	ctx := context.Background()

	t.Run("create cage invalid type", func(t *testing.T) {
		// Setup.
		input := v1.CreateCageRequest{
			Type:   "foobivore",
			Status: "ACTIVE",
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/cages", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "type")
	})

	t.Run("create cage invalid status", func(t *testing.T) {
		// Setup.
		input := v1.CreateCageRequest{
			Type:   "CARNIVORE",
			Status: "foobar",
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/cages", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "status")
	})
}

func TestUpdateCage(t *testing.T) {
	ctx := context.Background()

	t.Run("update cage invalid status", func(t *testing.T) {
		// Setup.
		input := v1.UpdateCageRequest{
			Status: "foobar",
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/v1/cages/1", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.UpdateCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "status")
	})
}

func TestAddDinoToCage(t *testing.T) {
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	cageID, dinoID := uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"id":     cageID.String(),
		"dinoId": dinoID.String(),
	})

	t.Run("add dino to cage powered down error", func(t *testing.T) {
		// Setup.
		ctrl := v1.Controller{
			Cage: cage.NewCore(&mockCageStore{
				getFunc: func() (cage.Cage, error) {
					return cage.Cage{
						ID:     cageID,
						Status: cage.CageStatusDown,
					}, nil
				},
			}, log, nil),
		}

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("/v1/cages/%s/dinosaurs/%s", cageID, dinoID), nil)
		require.NoError(t, err)

		// Execute.
		err = ctrl.AddDinosaurToCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		fmt.Println(tErr)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Equal(t, core.ErrInvalidCagePowerDown.Error(), tErr.Error())
	})

	t.Run("add dino to cage at capacity error", func(t *testing.T) {
		// Setup.
		ctrl := v1.Controller{
			Cage: cage.NewCore(&mockCageStore{
				getFunc: func() (cage.Cage, error) {
					return cage.Cage{
						ID:              cageID,
						Status:          cage.CageStatusActive,
						Capacity:        5,
						CurrentCapacity: 5,
					}, nil
				},
			}, log, nil),
		}

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("/v1/cages/%s/dinosaurs/%s", cageID, dinoID), nil)
		require.NoError(t, err)

		// Execute.
		err = ctrl.AddDinosaurToCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		fmt.Println(tErr)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Equal(t, core.ErrInvalidCageAtCapacity.Error(), tErr.Error())
	})

	t.Run("add dino to cage invalid type error", func(t *testing.T) {
		// Setup.
		ctrl := v1.Controller{
			Cage: cage.NewCore(&mockCageStore{
				getFunc: func() (cage.Cage, error) {
					return cage.Cage{
						ID:              cageID,
						Status:          cage.CageStatusActive,
						Capacity:        5,
						CurrentCapacity: 3,
						Type:            cage.CageTypeHerbivore,
					}, nil
				},
			},
				log,
				dino.NewCore(&mockDinoStore{
					getFunc: func() (dino.Dinosaur, error) {
						return dino.Dinosaur{
							ID:   dinoID,
							Diet: dino.DietTypeCarnivore,
						}, nil
					},
				}, log),
			),
		}

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("/v1/cages/%s/dinosaurs/%s", cageID, dinoID), nil)
		require.NoError(t, err)

		// Execute.
		err = ctrl.AddDinosaurToCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		fmt.Println(tErr)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Equal(t, core.ErrInvalidCageInvalidType.Error(), tErr.Error())
	})

	t.Run("add dino to cage different carnivore species error", func(t *testing.T) {
		// Setup.
		ctrl := v1.Controller{
			Cage: cage.NewCore(&mockCageStore{
				getFunc: func() (cage.Cage, error) {
					return cage.Cage{
						ID:              cageID,
						Status:          cage.CageStatusActive,
						Capacity:        5,
						CurrentCapacity: 3,
						Type:            cage.CageTypeCarnivore,
					}, nil
				},
			},
				log,
				dino.NewCore(&mockDinoStore{
					getFunc: func() (dino.Dinosaur, error) {
						return dino.Dinosaur{
							ID:      dinoID,
							Diet:    dino.DietTypeCarnivore,
							Species: dino.DinoSpeciesVelociraptor,
						}, nil
					},
					listByCageFunc: func() ([]dino.Dinosaur, error) {
						return []dino.Dinosaur{
							{
								ID:      dinoID,
								Diet:    dino.DietTypeCarnivore,
								Species: dino.DinoSpeciesTyrannosaurus,
							},
							{
								ID:      dinoID,
								Diet:    dino.DietTypeCarnivore,
								Species: dino.DinoSpeciesTyrannosaurus,
							},
						}, nil
					},
				}, log),
			),
		}

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("/v1/cages/%s/dinosaurs/%s", cageID, dinoID), nil)
		require.NoError(t, err)

		// Execute.
		err = ctrl.AddDinosaurToCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		fmt.Println(tErr)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Equal(t, core.ErrInvalidCageInvalidSpecies.Error(), tErr.Error())
	})
}

func TestRemoveDinoFromCage(t *testing.T) {
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	cageID, dinoID := uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"id":     cageID.String(),
		"dinoId": dinoID.String(),
	})

	t.Run("remove dino from cage invalid removal error", func(t *testing.T) {
		// Setup.
		ctrl := v1.Controller{
			Cage: cage.NewCore(&mockCageStore{
				getFunc: func() (cage.Cage, error) {
					return cage.Cage{
						ID:              cageID,
						CurrentCapacity: 0,
					}, nil
				},
			}, log, nil),
		}

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("/v1/cages/%s/dinosaurs/%s", cageID, dinoID), nil)
		require.NoError(t, err)

		// Execute.
		err = ctrl.RemoveDinosaurFromCage(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		fmt.Println(tErr)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Equal(t, core.ErrInvalidCageInvalidRemoval.Error(), tErr.Error())
	})
}
