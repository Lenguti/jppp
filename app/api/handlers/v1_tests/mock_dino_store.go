package v1_tests

import (
	"context"

	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/dino"
)

type mockDinoStore struct {
	dino.Storer

	getFunc        func() (dino.Dinosaur, error)
	listByCageFunc func() ([]dino.Dinosaur, error)
}

func (mds *mockDinoStore) Get(ctx context.Context, id string) (dino.Dinosaur, error) {
	return mds.getFunc()
}

func (mds *mockDinoStore) ListByCage(ctx context.Context, cageID string, filters ...core.Filter) ([]dino.Dinosaur, error) {
	return mds.listByCageFunc()
}
