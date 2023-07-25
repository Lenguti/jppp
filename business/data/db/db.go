package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DB - represents our data store orchestrator.
type DB struct {
	sql *sqlx.DB
	cfg Config
}

// New - returns an initialzed db.
func New(cfg Config) (*DB, error) {
	db, err := sqlx.Open("postgres", cfg.dbString())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return &DB{
		sql: db,
		cfg: cfg,
	}, nil
}

// Connect - helper methed for determining connectivity.
func (db *DB) Connect() error {
	if _, err := sqlx.Connect("postgres", db.cfg.dbString()); err != nil {
		return fmt.Errorf("connect: unable to connect to db: %w", err)
	}
	return nil
}

// Exec - execute db statements.
func (db *DB) Exec(ctx context.Context, query string, data any) error {
	if _, err := db.sql.NamedExecContext(ctx, query, data); err != nil {
		return fmt.Errorf("exec: unable to named exec: %w", err)
	}
	return nil
}

// Get - fetch db item.
func (db *DB) Get(ctx context.Context, data any, query string, val string) error {
	return db.sql.GetContext(ctx, data, query, val)
}

// List - list db items.
func (db *DB) List(ctx context.Context, data any, query string, vals ...string) error {
	var ivals []any
	for i := range vals {
		ivals = append(ivals, vals[i])
	}
	return db.sql.SelectContext(ctx, data, query, ivals...)
}

// BeginTX - starts a db transaction.
func (db *DB) BeginTx(ctx context.Context) *sqlx.Tx {
	return db.sql.MustBeginTx(ctx, nil)
}

// CommitTx - commits a db transaction.
func (db *DB) CommitTx(tx *sqlx.Tx) error {
	return tx.Commit()
}
