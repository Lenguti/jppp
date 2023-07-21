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
