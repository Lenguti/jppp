package dino

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core"
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

// Get - will featch a dino by its id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Dinosaur, error) {
	d, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Dinosaur{}, fmt.Errorf("get: failed to fetch dino: %w", err)
	}
	return d, nil
}

// UpdateName - will update the name of the provided dino.
func (c *Core) UpdateName(ctx context.Context, id uuid.UUID, name string) (Dinosaur, error) {
	d, err := c.Get(ctx, id)
	if err != nil {
		return Dinosaur{}, fmt.Errorf("update name: unable to fetch dinosaur: %w", err)
	}

	if d.Name == name {
		return d, nil
	}

	now := time.Now().UTC()
	d.Name = name
	d.UpdatedAt = now
	if err := c.store.UpdateName(ctx, d.ID.String(), d.Name, d.UpdatedAt); err != nil {
		return Dinosaur{}, fmt.Errorf("update status: failed to update dino: %w", err)
	}

	return d, nil
}

// ListByCageID - will list all dinos for a given cage.
func (c *Core) ListByCageID(ctx context.Context, cageID uuid.UUID, filters ...core.Filter) ([]Dinosaur, error) {
	c.log.Info().Fields(map[string]any{"filters": filters}).Msg("Listing Dinos in cage.")
	dinos, err := c.store.ListByCage(ctx, cageID.String(), filters...)
	if err != nil {
		return nil, fmt.Errorf("list by cage: failed to list dinos: %w", err)
	}
	return dinos, nil
}

// List - will list all dinosaurs.
func (c *Core) List(ctx context.Context) ([]Dinosaur, error) {
	ds, err := c.store.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list: failed to list dinos: %w", err)
	}
	return ds, nil
}
