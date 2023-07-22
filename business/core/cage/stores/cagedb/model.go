package cagedb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core/cage"
)

type dbCage struct {
	ID              string `db:"id"`
	Type            string `db:"type"`
	Capacity        int    `db:"capacity"`
	CurrentCapacity int    `db:"current_capacity"`
	Status          string `db:"status"`
	CreatedAt       int64  `db:"created_at"`
	UpdateAt        int64  `db:"updated_at"`
}

func toDBCage(c cage.Cage) dbCage {
	return dbCage{
		ID:              c.ID.String(),
		Type:            c.Type.String(),
		Capacity:        c.Capacity,
		CurrentCapacity: c.CurrentCapacity,
		Status:          c.Status.String(),
		CreatedAt:       c.CreatedAt.Unix(),
		UpdateAt:        c.UpdatedAt.Unix(),
	}
}

func toCoreCages(dbcages []dbCage) []cage.Cage {
	cages := make([]cage.Cage, 0, len(dbcages))
	for _, v := range dbcages {
		cages = append(cages, toCoreCage(v))
	}
	return cages
}

func toCoreCage(dbc dbCage) cage.Cage {
	return cage.Cage{
		ID:              uuid.MustParse(dbc.ID),
		Type:            cage.Type(dbc.Type),
		Capacity:        dbc.Capacity,
		CurrentCapacity: dbc.CurrentCapacity,
		Status:          cage.Status(dbc.Status),
		CreatedAt:       time.Unix(dbc.CreatedAt, 0),
		UpdatedAt:       time.Unix(dbc.UpdateAt, 0),
	}
}
