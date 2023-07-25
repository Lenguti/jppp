package cage

import (
	"context"
	"time"

	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for cages.
type Storer interface {
	Create(ctx context.Context, c Cage) error
	Get(ctx context.Context, id string) (Cage, error)
	List(ctx context.Context, filters ...core.Filter) ([]Cage, error)
	UpdateStatus(ctx context.Context, id, status string, ts time.Time) error
	AddDino(ctx context.Context, c Cage, dinoID string) error
	RemoveDino(ctx context.Context, c Cage, dinoID string) error
}

// Core - represents the core business logic for cages.
type Core struct {
	store Storer
	log   zerolog.Logger
	dino  *dino.Core
}

// NewCore - returns a new cage core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger, dc *dino.Core) *Core {
	return &Core{
		store: store,
		log:   log,
		dino:  dc,
	}
}
