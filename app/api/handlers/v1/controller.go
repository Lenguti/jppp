package v1

import (
	"fmt"

	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/core/cage/stores/cagedb"
	"github.com/lenguti/jppp/business/data/db"
	"github.com/lenguti/jppp/foundation/api"
	"github.com/rs/zerolog"
)

type Controller struct {
	Config Config
	Log    zerolog.Logger
	Router *api.Router
	Cage   *cage.Core

	db *db.DB
}

func NewController(log zerolog.Logger, cfg Config) (*Controller, error) {
	ddb, err := db.New(db.Config{
		User:         cfg.DBUser,
		Password:     cfg.DBPass,
		Name:         cfg.DBName,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("new controller: unable to initialize new cage core: %w", err)
	}

	return &Controller{
		Config: cfg,
		Log:    log,
		Router: api.NewRouter(),
		Cage:   cage.NewCore(cagedb.NewStore(ddb)),

		db: ddb,
	}, nil
}
