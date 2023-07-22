package dinodb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core/dino"
)

type dbDino struct {
	ID        string  `db:"id"`
	CageID    *string `db:"cage_id"`
	Name      string  `db:"name"`
	Species   string  `db:"species"`
	Diet      string  `db:"diet"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func toDBDino(d dino.Dinosaur) dbDino {
	dbd := dbDino{
		ID:        d.ID.String(),
		Name:      d.Name,
		Species:   d.Species,
		Diet:      d.Diet.String(),
		CreatedAt: d.CreatedAt.Unix(),
		UpdatedAt: d.UpdatedAt.Unix(),
	}
	if d.CageID != uuid.Nil {
		dbd.CageID = toStrPtr(d.CageID.String())
	}
	return dbd
}

func toCoreDinos(dbDinos []dbDino) []dino.Dinosaur {
	dinos := make([]dino.Dinosaur, 0, len(dbDinos))
	for _, v := range dbDinos {
		dinos = append(dinos, toCoreDino(v))
	}
	return dinos
}

func toCoreDino(dbd dbDino) dino.Dinosaur {
	d := dino.Dinosaur{
		ID:        uuid.MustParse(dbd.ID),
		Name:      dbd.Name,
		Species:   dbd.Species,
		Diet:      dino.Diet(dbd.Diet),
		CreatedAt: time.Unix(dbd.CreatedAt, 0),
		UpdatedAt: time.Unix(dbd.UpdatedAt, 0),
	}
	if dbd.CageID != nil {
		d.CageID = uuid.MustParse(*dbd.CageID)
	}
	return d
}

func toStrPtr(v string) *string {
	return &v
}
