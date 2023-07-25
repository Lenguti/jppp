package cagedb

import (
	"testing"

	"github.com/lenguti/jppp/business/core"
	"github.com/stretchr/testify/assert"
)

func TestListClauseBuilder(t *testing.T) {
	t.Run("no filters", func(t *testing.T) {
		want := `
	SELECT *
	FROM cage
	`
		var wantVals []string
		got, gotVals := listClauseBuilder()
		assert.Equal(t, want, got)
		assert.Equal(t, wantVals, gotVals)
	})

	t.Run("status filter", func(t *testing.T) {
		want := `
	SELECT *
	FROM cage
	WHERE status = $1`

		wantVals := []string{"ACTIVE"}
		got, gotVals := listClauseBuilder(core.Filter{Key: "status", Value: "ACTIVE"})
		assert.Equal(t, want, got)
		assert.Equal(t, wantVals, gotVals)
	})
}
