package v1_tests

import (
	"context"

	"github.com/lenguti/jppp/business/core/cage"
)

type mockCageStore struct {
	cage.Storer

	getFunc func() (cage.Cage, error)
}

func (mcs *mockCageStore) Get(ctx context.Context, id string) (cage.Cage, error) {
	return mcs.getFunc()
}
