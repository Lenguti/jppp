package dino

import "fmt"

const (
	DinoSpeciesTyrannosaurus = "Tyrannosaurus"
	DinoSpeciesVelociraptor  = "Velociraptor"
	DinoSpeciesSpinosaurus   = "Spinosaurus"
	DinoSpeciesMegalosaurus  = "Megalosaurus"

	DinoSpeciesBrachiosaurus = "Brachiosaurus"
	DinoSpeciesStegosaurus   = "Stegosaurus"
	DinoSpeciesAnkylosaurus  = "Ankylosaurus"
	DinoSpeciesTriceratops   = "Triceratops"
)

var validDinoSpecies = map[string]struct{}{
	DinoSpeciesTyrannosaurus: {},
	DinoSpeciesVelociraptor:  {},
	DinoSpeciesSpinosaurus:   {},
	DinoSpeciesMegalosaurus:  {},
	DinoSpeciesBrachiosaurus: {},
	DinoSpeciesStegosaurus:   {},
	DinoSpeciesAnkylosaurus:  {},
	DinoSpeciesTriceratops:   {},
}

func ParseSpecies(v string) error {
	if _, ok := validDinoSpecies[v]; !ok {
		return fmt.Errorf("parse species: invalid dino species")
	}
	return nil
}

type Diet string

func (d Diet) String() string {
	return string(d)
}

const (
	DietTypeCarnivore = "CARNIVORE"
	DietTypeHerbivore = "HERBIVORE"
)

var validDietTypes = map[Diet]struct{}{
	DietTypeCarnivore: {},
	DietTypeHerbivore: {},
}

func ParseDiet(v string) error {
	if _, ok := validDietTypes[Diet(v)]; !ok {
		return fmt.Errorf("parse diet: invalid diet type")
	}
	return nil
}

var DinoSpeciesMapping = map[string]Diet{
	DinoSpeciesTyrannosaurus: DietTypeCarnivore,
	DinoSpeciesVelociraptor:  DietTypeCarnivore,
	DinoSpeciesSpinosaurus:   DietTypeCarnivore,
	DinoSpeciesMegalosaurus:  DietTypeCarnivore,
	DinoSpeciesBrachiosaurus: DietTypeHerbivore,
	DinoSpeciesStegosaurus:   DietTypeHerbivore,
	DinoSpeciesAnkylosaurus:  DietTypeHerbivore,
	DinoSpeciesTriceratops:   DietTypeHerbivore,
}
