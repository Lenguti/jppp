package dinodb

import (
	"context"
	"fmt"
	"strings"

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
	q := listClauseBuilder()
	var out []dbDino
	if err := s.db.List(ctx, &out, q); err != nil {
		return nil, fmt.Errorf("list: failed to list dinos: %w", err)
	}
	return toCoreDinos(out), nil
}

// ListByCage - will fetch all dinos associated to the provided cage ids.
func (s *Store) ListByCage(ctx context.Context, ids ...string) ([]dino.Dinosaur, error) {
	q := listClauseBuilder(ids...)
	var out []dbDino
	if err := s.db.List(ctx, &out, q, ids...); err != nil {
		return nil, fmt.Errorf("list by cage: failed to list dinos: %w", err)
	}
	return toCoreDinos(out), nil
}

func listClauseBuilder(ids ...string) string {
	const q = `
	SELECT *
	FROM dinosaur
	`

	var b strings.Builder
	b.WriteString(q)
	if len(ids) == 0 {
		return b.String()
	}

	const whereClause = "WHERE cage_id = $1"
	if len(ids) == 1 {
		b.WriteString(whereClause)
		return b.String()
	}

	const inClause = "WHERE cage_id IN ("
	b.WriteString(inClause)
	for i := range ids {
		b.WriteString(fmt.Sprintf("$%d", i+1))
		if i != len(ids)-1 {
			b.WriteString(", ")

		} else {
			b.WriteString(")")
		}
	}
	return b.String()
}
