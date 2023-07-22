package v1

import (
	"github.com/lenguti/jppp/business/core/cage"
	"github.com/lenguti/jppp/business/core/dino"
)

// ClientCage - represents a client cage entity.
type ClientCage struct {
	ID              string       `json:"id"`
	Type            string       `json:"type"`
	Capacity        int          `json:"capacity"`
	CurrentCapacity int          `json:"currentCapacity"`
	Status          string       `json:"status"`
	Dinosaurs       []ClientDino `json:"dinosaurs"`
	CreatedAt       int64        `json:"createdAt"`
	UpdatedAt       int64        `json:"updatedAt"`
}

func toCoreNewCage(input CreateCageRequest) cage.NewCage {
	newCage := cage.NewCage{
		Type:     cage.Type(input.Type),
		Capacity: input.Capacity,
		Status:   cage.Status(input.Status),
	}
	return newCage
}

func toClientCages(cages []cage.Cage, dinos []dino.Dinosaur) []ClientCage {
	ccages := make([]ClientCage, 0, len(cages))

	cageDinoMap := map[string][]dino.Dinosaur{}
	for _, v := range dinos {
		if _, ok := cageDinoMap[v.CageID.String()]; !ok {
			cageDinoMap[v.CageID.String()] = []dino.Dinosaur{v}
			continue
		}
		cageDinoMap[v.CageID.String()] = append(cageDinoMap[v.CageID.String()], v)
	}

	for _, cage := range cages {
		ccages = append(ccages, toClientCage(cage, cageDinoMap[cage.ID.String()]))
	}
	return ccages
}

func toClientCage(input cage.Cage, dinos []dino.Dinosaur) ClientCage {
	return ClientCage{
		ID:              input.ID.String(),
		Type:            input.Type.String(),
		Capacity:        input.Capacity,
		CurrentCapacity: input.CurrentCapacity,
		Status:          input.Status.String(),
		Dinosaurs:       toClientDinos(dinos),
		CreatedAt:       input.CreatedAt.Unix(),
		UpdatedAt:       input.UpdatedAt.Unix(),
	}
}
