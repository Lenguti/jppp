package v1

import (
	"strings"

	"github.com/google/uuid"
	"github.com/lenguti/jppp/business/core/dino"
)

type ClientDino struct {
	ID        string `json:"id"`
	CageID    string `json:"cage_id,omitempty"`
	Name      string `json:"name"`
	Species   string `json:"species"`
	Diet      string `json:"diet"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func toCoreNewDino(input CreateDinoRequest) dino.NewDino {
	newDino := dino.NewDino{
		Name:    input.Name,
		Species: strings.Title(input.Species),
		Diet:    dino.Diet(strings.ToUpper(input.Diet)),
	}
	return newDino
}

type ClientDinoSpecies struct {
	Species string `json:"species"`
	Diet    string `json:"diet"`
}

func toClientDinos(input []dino.Dinosaur) []ClientDino {
	dinos := make([]ClientDino, 0, len(input))
	for _, v := range input {
		dinos = append(dinos, toClientDino(v))
	}
	return dinos
}

func toClientDino(input dino.Dinosaur) ClientDino {
	cd := ClientDino{
		ID:        input.ID.String(),
		Name:      input.Name,
		Species:   input.Species,
		Diet:      input.Diet.String(),
		CreatedAt: input.CreatedAt.Unix(),
		UpdatedAt: input.UpdatedAt.Unix(),
	}
	if input.CageID != uuid.Nil {
		cd.CageID = input.CageID.String()
	}
	return cd
}
