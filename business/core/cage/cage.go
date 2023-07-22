package cage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new cage.
func (c *Core) Create(ctx context.Context, nc NewCage) (Cage, error) {
	now := time.Now().UTC()
	cg := Cage{
		ID:              uuid.New(),
		Type:            nc.Type,
		Capacity:        nc.Capacity,
		CurrentCapacity: 0,
		Status:          nc.Status,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if err := c.store.Create(ctx, cg); err != nil {
		return Cage{}, fmt.Errorf("create: failed to create cage: %w", err)
	}

	return cg, nil
}

// Get - will featch a cage by its id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Cage, error) {
	cg, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Cage{}, fmt.Errorf("create: failed to fetch cage: %w", err)
	}
	return cg, nil
}

// List - will list all cages.
func (c *Core) List(ctx context.Context) ([]Cage, error) {
	cgs, err := c.store.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list: failed to list cages: %w", err)
	}
	return cgs, nil
}
