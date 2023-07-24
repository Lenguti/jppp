package dinodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lenguti/jppp/business/core"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/lenguti/jppp/business/data/db"
)

// Store - manages the set of apis for dino database access.
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
func (s *Store) Create(ctx context.Context, d dino.Dinosaur) error {
	dbDino := toDBDino(d)
	const q = `
	INSERT INTO dinosaur (
		id,
		cage_id,
		name,
		species,
		diet,
		created_at,
		updated_at
	) VALUES (
		:id,
		:cage_id,
		:name,
		:species,
		:diet,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbDino); err != nil {
		return fmt.Errorf("create: failed to create dino: %w", err)
	}
	return nil
}

// Get - will fetch a dino by its id.
func (s *Store) Get(ctx context.Context, id string) (dino.Dinosaur, error) {
	const q = `
	SELECT *
	FROM dinosaur
	WHERE id = $1
	`
	var out dbDino
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		return dino.Dinosaur{}, fmt.Errorf("get: failed to fetch dino: %w", err)
	}
	return toCoreDino(out), nil
}

// List - will list all dinos.
func (s *Store) List(ctx context.Context) ([]dino.Dinosaur, error) {
	const q = `
	SELECT *
	FROM dinosaur
	`
	var out []dbDino
	if err := s.db.List(ctx, &out, q); err != nil {
		return nil, fmt.Errorf("list: failed to list dinos: %w", err)
	}
	return toCoreDinos(out), nil
}

// UpdateName - will update the name of a dino.
func (s *Store) UpdateName(ctx context.Context, id, name string, ts time.Time) error {
	const q = `
	UPDATE dinosaur
	SET
	name = :name,
	updated_at = :updated_at
	WHERE id = :id
	`
	if err := s.db.Exec(ctx, q, map[string]any{"name": name, "updated_at": ts.Unix(), "id": id}); err != nil {
		return fmt.Errorf("get: failed to update dino name: %w", err)
	}
	return nil
}

// ListByCage - will fetch all dinos associated to the provided cage ids.
func (s *Store) ListByCage(ctx context.Context, cageID string, filters ...core.Filter) ([]dino.Dinosaur, error) {
	q, vals := listClauseBuilder(cageID, filters...)
	var out []dbDino
	if err := s.db.List(ctx, &out, q, vals...); err != nil {
		return nil, fmt.Errorf("list by cage: failed to list dinos: %w", err)
	}
	return toCoreDinos(out), nil
}

func listClauseBuilder(cageID string, filters ...core.Filter) (string, []string) {
	const q = `
	SELECT *
	FROM dinosaur
	WHERE cage_id = $1
	`
	vals := []string{cageID}
	if len(filters) == 0 {
		return q, vals
	}

	filterMap := map[string]string{
		"species": "species = $%d",
	}

	var b strings.Builder
	b.WriteString(q)
	b.WriteString("AND ")
	for i := 0; i < len(filters); i++ {
		c, ok := filterMap[filters[i].Key]
		if ok {
			b.WriteString(fmt.Sprintf(c, i+2))
			vals = append(vals, filters[i].Value)
		}
	}
	return b.String(), vals
}
