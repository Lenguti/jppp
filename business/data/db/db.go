package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	sql *sqlx.DB
	cfg Config
}

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

func (db *DB) Connect() error {
	if _, err := sqlx.Connect("postgres", db.cfg.dbString()); err != nil {
		return fmt.Errorf("connect: unable to connect to db: %w", err)
	}
	return nil
}

func (db *DB) Exec(ctx context.Context, query string, data any) error {
	if _, err := db.sql.NamedExecContext(ctx, query, data); err != nil {
		return fmt.Errorf("exec: unable to named exec: %w", err)
	}
	return nil
}

func (db *DB) Get(ctx context.Context, data any, query string, val string) error {
	return db.sql.GetContext(ctx, data, query, val)
}

func (db *DB) List(ctx context.Context, data any, query string, vals ...string) error {
	var v any
	switch len(vals) {
	case 0:
		return db.sql.SelectContext(ctx, data, query)
	case 1:
		v = vals[0]
	default:
		v = vals
	}
	return db.sql.SelectContext(ctx, data, query, v)
}
