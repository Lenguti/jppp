package dino

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for dinos.
type Storer interface {
	Create(ctx context.Context, d Dinosaur) error
	ListByCage(ctx context.Context, id ...string) ([]Dinosaur, error)
	Get(ctx context.Context, id string) (Dinosaur, error)
	List(ctx context.Context) ([]Dinosaur, error)
}

// Core - represents the core business logic for dinos.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new dino core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
