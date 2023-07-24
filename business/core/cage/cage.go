package cage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core"
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
		return Cage{}, fmt.Errorf("get: failed to fetch cage: %w", err)
	}
	return cg, nil
}

// List - will list all cages.
func (c *Core) List(ctx context.Context, filters ...core.Filter) ([]Cage, error) {
	c.log.Info().Fields(map[string]any{"filters": filters}).Msg("Listing cages.")
	cgs, err := c.store.List(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("list: failed to list cages: %w", err)
	}
	return cgs, nil
}

// UpdateStatus - will update the status of the provided cage.
func (c *Core) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) (Cage, error) {
	cge, err := c.Get(ctx, id)
	if err != nil {
		return Cage{}, fmt.Errorf("update status: unable to fetch cage: %w", err)
	}

	if cge.Status == status {
		return cge, nil
	}

	if status == CageStatusDown && cge.CurrentCapacity > 0 {
		return Cage{}, core.ErrPowerDownCage
	}

	now := time.Now().UTC()
	cge.Status = status
	cge.UpdatedAt = now
	if err := c.store.UpdateStatus(ctx, cge.ID.String(), cge.Status.String(), cge.UpdatedAt); err != nil {
		return Cage{}, fmt.Errorf("update status: failed to update cage: %w", err)
	}

	return cge, nil
}
