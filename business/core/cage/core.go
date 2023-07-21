package cage

import (
	"context"
)

// Storer - represents the data layer behavior for cages.
type Storer interface {
	Create(ctx context.Context, c Cage) error
}

// Core - represents the core business logic for cages.
type Core struct {
	store Storer
}

// NewCore - returns a new cage core with all its components initialized.
func NewCore(store Storer) *Core {
	return &Core{
		store: store,
	}
}
