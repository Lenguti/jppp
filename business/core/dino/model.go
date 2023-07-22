package dino

import (
	"time"

	"github.com/google/uuid"
)

// Dinosaur - represents a business domain dinosaur.
type Dinosaur struct {
	ID        uuid.UUID
	CageID    uuid.UUID
	Name      string
	Species   string
	Diet      Diet
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewDino - represents fields needed to create a new dinosaur.
type NewDino struct {
	Name    string
	Species string
	Diet    Diet
}
