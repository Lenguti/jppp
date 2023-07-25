package v1_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/lenguti/jppp/app/api/handlers/v1"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/lenguti/jppp/foundation/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDino(t *testing.T) {
	ctx := context.Background()

	t.Run("create dino invalid name", func(t *testing.T) {
		// Setup.
		input := v1.CreateDinoRequest{
			Name:    "",
			Species: dino.DinoSpeciesSpinosaurus,
			Diet:    dino.DietTypeCarnivore,
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/dinosaurs", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateDino(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "name")
	})

	t.Run("create dino invalid species", func(t *testing.T) {
		// Setup.
		input := v1.CreateDinoRequest{
			Name:    "Gerber",
			Species: "Gorilla",
			Diet:    dino.DietTypeCarnivore,
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/dinosaurs", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateDino(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "species")
	})

	t.Run("create dino invalid diet", func(t *testing.T) {
		// Setup.
		input := v1.CreateDinoRequest{
			Name:    "Gerber",
			Species: dino.DinoSpeciesAnkylosaurus,
			Diet:    "Fruitivore",
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/dinosaurs", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateDino(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "diet")
	})

	t.Run("create dino invalid species diet", func(t *testing.T) {
		// Setup.
		input := v1.CreateDinoRequest{
			Name:    "Gerber",
			Species: dino.DinoSpeciesAnkylosaurus,
			Diet:    dino.DietTypeCarnivore,
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/dinosaurs", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateDino(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "species diet")
	})
}

func TestUpdateDino(t *testing.T) {
	ctx := context.Background()

	t.Run("update dino invalid name", func(t *testing.T) {
		// Setup.
		input := v1.UpdateDinoRequest{
			Name: "",
		}
		ctrl := v1.Controller{}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/v1/dinosaurs/1", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.UpdateDino(ctx, w, r)

		// Validate.
		require.Error(t, err)
		tErr, ok := err.(api.HTTPError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, tErr.Err.StatusCode)
		assert.Contains(t, tErr.Err.Details, "name")
	})
}
