package cagedb

import (
	"context"
	"fmt"

	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/data/db"
)

// Store - manages the set of apis for cage database access.
type Store struct {
	db *db.DB
}

// NewStore - constructs the api for data access.
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

// Create - will insert a new cage record.
func (s *Store) Create(ctx context.Context, c cage.Cage) error {
	dbCage := toDBCage(c)
	const q = `
	INSERT INTO cage (
		id,
		type,
		capacity,
		current_capacity,
		status,
		created_at,
		updated_at
	) VALUES (
		:id,
		:type,
		:capacity,
		:current_capacity,
		:status,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbCage); err != nil {
		return fmt.Errorf("create: failed to create cage: %w", err)
	}

	return nil
}
