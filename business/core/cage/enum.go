package cage

import (
	"fmt"
	"strings"
)

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	CageTypeHerbivore = "HERBIVORE"
	CageTypeCarnivore = "CARNIVORE"
)

var validCageType = map[Type]struct{}{
	CageTypeHerbivore: {},
	CageTypeCarnivore: {},
}

func ParseType(v string) error {
	if _, ok := validCageType[Type(strings.ToUpper(v))]; !ok {
		return fmt.Errorf("parse type: invalid cage type")
	}
	return nil
}

type Status string

func (s Status) String() string {
	return string(s)
}

const (
	CageStatusActive = "ACTIVE"
	CageStatusDown   = "DOWN"
)

var validCageStatus = map[Status]struct{}{
	CageStatusActive: {},
	CageStatusDown:   {},
}

func ParseStatus(v string) error {
	if _, ok := validCageStatus[Status(strings.ToUpper(v))]; !ok {
		return fmt.Errorf("parse status: invalid cage status")
	}
	return nil
}
