package v1

import (
	"github.com/lenguti/jppp/business/core/cage"
)

// ClientCage - represents a client cage entity.
type ClientCage struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Capacity        int    `json:"capacity"`
	CurrentCapacity int    `json:"currentCapacity"`
	Status          string `json:"status"`
	CreatedAt       int64  `json:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt"`
}

func toCoreNewCage(input CreateCageRequest) cage.NewCage {
	newCage := cage.NewCage{
		Type:     cage.Type(input.Type),
		Capacity: input.Capacity,
		Status:   cage.Status(input.Status),
	}
	return newCage
}

func toClientCage(input cage.Cage) ClientCage {
	return ClientCage{
		ID:              input.ID.String(),
		Type:            input.Type.String(),
		Capacity:        input.Capacity,
		CurrentCapacity: input.CurrentCapacity,
		Status:          input.Status.String(),
		CreatedAt:       input.CreatedAt.Unix(),
		UpdatedAt:       input.UpdatedAt.Unix(),
	}
}
