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

func toCoreCage(dbCage dbCage) cage.Cage {
	return cage.Cage{
		ID:              uuid.MustParse(dbCage.ID),
		Type:            cage.Type(dbCage.Type),
		Capacity:        dbCage.Capacity,
		CurrentCapacity: dbCage.CurrentCapacity,
		Status:          cage.Status(dbCage.Status),
		CreatedAt:       time.Unix(dbCage.CreatedAt, 0),
		UpdatedAt:       time.Unix(dbCage.UpdateAt, 0),
	}
}
