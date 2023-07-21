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
		return err
	}
	return nil
}
