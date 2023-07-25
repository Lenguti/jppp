package v1

import (
	"fmt"

	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/core/cage/stores/cagedb"
	"github.com/lenguti/jppp/business/core/dino"
	"github.com/lenguti/jppp/business/core/dino/stores/dinodb"
	"github.com/lenguti/jppp/business/data/db"
	"github.com/lenguti/jppp/foundation/api"
	"github.com/rs/zerolog"
)

type Controller struct {
	Cage *cage.Core
	Dino *dino.Core

	db     *db.DB
	config Config
	log    zerolog.Logger
	router *api.Router
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
		return nil, fmt.Errorf("new controller: unable to initialize new db: %w", err)
	}

	dc := dino.NewCore(dinodb.NewStore(ddb), log)
	cc := cage.NewCore(cagedb.NewStore(ddb), log, dc)

	return &Controller{
		Cage: cc,
		Dino: dc,

		db:     ddb,
		config: cfg,
		log:    log,
		router: api.NewRouter(),
	}, nil
}
