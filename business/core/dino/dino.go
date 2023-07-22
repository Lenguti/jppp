package dino

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new dino.
func (c *Core) Create(ctx context.Context, nd NewDino) (Dinosaur, error) {
	now := time.Now().UTC()
	d := Dinosaur{
		ID:        uuid.New(),
		CageID:    uuid.Nil,
		Name:      nd.Name,
		Species:   nd.Species,
		Diet:      nd.Diet,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := c.store.Create(ctx, d); err != nil {
		return Dinosaur{}, fmt.Errorf("create: failed to create dino: %w", err)
	}

	return d, nil
}

// ListByCageID - will list all dinos for a given cage.
func (c *Core) ListByCageID(ctx context.Context, cageID uuid.UUID) ([]Dinosaur, error) {
	dinos, err := c.store.ListByCage(ctx, cageID.String())
	if err != nil {
		return nil, fmt.Errorf("list by cage: failed to list dinos: %w", err)
	}
	return dinos, nil
}

// ListByCageIDs - will list all dinos for the provided cages.
func (c *Core) ListByCageIDs(ctx context.Context, cageIDs ...uuid.UUID) ([]Dinosaur, error) {
	cageIDStrs := make([]string, 0, len(cageIDs))
	for _, v := range cageIDs {
		cageIDStrs = append(cageIDStrs, v.String())
	}
	dinos, err := c.store.ListByCage(ctx, cageIDStrs...)
	if err != nil {
		return nil, fmt.Errorf("list by cage: failed to list dinos: %w", err)
	}
	return dinos, nil
}
