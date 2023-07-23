package cage

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for cages.
type Storer interface {
	Create(ctx context.Context, c Cage) error
	Get(ctx context.Context, id string) (Cage, error)
	List(ctx context.Context, filters ...Filter) ([]Cage, error)
	UpdateStatus(ctx context.Context, id, status string, ts time.Time) error
}

// Core - represents the core business logic for cages.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new cage core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
