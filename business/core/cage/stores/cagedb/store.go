package cagedb

import (
	"context"
	"fmt"
	"strings"
	"time"

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

// UpdateStatus - will update the status of a cage.
func (s *Store) UpdateStatus(ctx context.Context, id, status string, ts time.Time) error {
	const q = `
	UPDATE cage
	SET
	status = :status,
	updated_at = :updated_at
	WHERE id = :id
	`
	if err := s.db.Exec(ctx, q, map[string]any{"status": status, "updated_at": ts.Unix(), "id": id}); err != nil {
		return fmt.Errorf("get: failed to update cage status: %w", err)
	}
	return nil
}

// Get - will fetch a cage by its id.
func (s *Store) Get(ctx context.Context, id string) (cage.Cage, error) {
	const q = `
	SELECT *
	FROM cage
	WHERE id = $1
	`
	var out dbCage
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		return cage.Cage{}, fmt.Errorf("get: failed to fetch cage: %w", err)
	}
	return toCoreCage(out), nil
}

// List - will list all cages.
func (s *Store) List(ctx context.Context, filters ...cage.Filter) ([]cage.Cage, error) {
	q, vals := listClauseBuilder(filters...)
	var out []dbCage
	if err := s.db.List(ctx, &out, q, vals...); err != nil {
		return nil, fmt.Errorf("list: failed to list cages: %w", err)
	}
	return toCoreCages(out), nil
}

func listClauseBuilder(filters ...cage.Filter) (string, []string) {
	const q = `
	SELECT *
	FROM cage
	`

	if len(filters) == 0 {
		return q, nil
	}

	filterMap := map[string]string{
		"status": "status = $%d",
	}

	vals := make([]string, 0, len(filters))
	var b strings.Builder
	b.WriteString(q)
	b.WriteString("WHERE ")
	for i := 0; i < len(filters); i++ {
		c, ok := filterMap[filters[i].Key]
		if ok {
			b.WriteString(fmt.Sprintf(c, i+1))
			vals = append(vals, filters[i].Value)
		}
		if i != len(filters)-1 {
			b.WriteString(" AND ")
		}
	}
	return b.String(), vals
}
