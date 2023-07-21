package cage

import (
	"time"

	"github.com/google/uuid"
)

// Cage - represents a business domain cage.
type Cage struct {
	ID              uuid.UUID
	Type            Type
	Capacity        int
	CurrentCapacity int
	Status          Status
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// NewCage - represents fields needed to create a new cage.
type NewCage struct {
	Type     Type
	Capacity int
	Status   Status
}
