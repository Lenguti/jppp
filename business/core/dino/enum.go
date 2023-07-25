package dino

import (
	"fmt"
	"strings"
)

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

var validDinoSpecies = map[string]Diet{
	DinoSpeciesTyrannosaurus: DietTypeCarnivore,
	DinoSpeciesVelociraptor:  DietTypeCarnivore,
	DinoSpeciesSpinosaurus:   DietTypeCarnivore,
	DinoSpeciesMegalosaurus:  DietTypeCarnivore,
	DinoSpeciesBrachiosaurus: DietTypeHerbivore,
	DinoSpeciesStegosaurus:   DietTypeHerbivore,
	DinoSpeciesAnkylosaurus:  DietTypeHerbivore,
	DinoSpeciesTriceratops:   DietTypeHerbivore,
}

// ParseSpecies - will attempt to validate the provided species and return their diet.
func ParseSpecies(v string) (Diet, error) {
	d, ok := validDinoSpecies[strings.Title(v)]
	if !ok {
		return "", fmt.Errorf("parse species: invalid dino species")
	}
	return d, nil
}

// Diet - represents dino diet enum.
type Diet string

// String - returns string representation of diet.
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

// ParseDiet - will attempt to validate the provided diet.
func ParseDiet(v string) error {
	if _, ok := validDietTypes[Diet(strings.ToUpper(v))]; !ok {
		return fmt.Errorf("parse diet: invalid diet type")
	}
	return nil
}

// DinoSpeciesMapping - mapping of available species and their diets.
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
