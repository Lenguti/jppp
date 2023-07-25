package dinodb

import (
	"testing"

	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/stretchr/testify/assert"
)

func TestListClauseBuilder(t *testing.T) {
	t.Run("no filters", func(t *testing.T) {
		want := `
	SELECT *
	FROM dinosaur
	WHERE cage_id = $1
	`
		cageID := "foo"
		wantVals := []string{cageID}
		got, gotVals := listClauseBuilder(cageID)
		assert.Equal(t, want, got)
		assert.Equal(t, wantVals, gotVals)
	})

	t.Run("species filter", func(t *testing.T) {
		want := `
	SELECT *
	FROM dinosaur
	WHERE cage_id = $1
	AND species = $2`
		cageID := "foo"
		wantVals := []string{cageID, dino.DinoSpeciesAnkylosaurus}
		got, gotVals := listClauseBuilder(cageID, core.Filter{Key: "species", Value: dino.DinoSpeciesAnkylosaurus})
		assert.Equal(t, want, got)
		assert.Equal(t, wantVals, gotVals)
	})
}
