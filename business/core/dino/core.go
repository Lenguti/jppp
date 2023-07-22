package dino

import (
	"context"
)

// Storer - represents the data layer behavior for dinos.
type Storer interface {
	Create(ctx context.Context, d Dinosaur) error
	ListByCage(ctx context.Context, id ...string) ([]Dinosaur, error)
}

// Core - represents the core business logic for dinos.
type Core struct {
	store Storer
}

// NewCore - returns a new dino core with all its components initialized.
func NewCore(store Storer) *Core {
	return &Core{
		store: store,
	}
}
