package dino

import (
	"context"
	"time"

	"github.com/lenguti/jppp/business/core"
	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for dinos.
type Storer interface {
	Create(ctx context.Context, d Dinosaur) error
	ListByCage(ctx context.Context, cageID string, filters ...core.Filter) ([]Dinosaur, error)
	Get(ctx context.Context, id string) (Dinosaur, error)
	List(ctx context.Context) ([]Dinosaur, error)
	UpdateName(ctx context.Context, id, name string, ts time.Time) error
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
