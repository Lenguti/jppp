package cage

import (
	"fmt"
	"strings"
)

// Type - represents cage type enum.
type Type string

// String - returns string representation of type.
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

// ParseType - will attempt to validate the provided type.
func ParseType(v string) error {
	if _, ok := validCageType[Type(strings.ToUpper(v))]; !ok {
		return fmt.Errorf("parse type: invalid cage type")
	}
	return nil
}

// Status - represents cage status enum.
type Status string

// String - returns string representation of status.
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

// ParseType - will attempt to validate the provided status.
func ParseStatus(v string) error {
	if _, ok := validCageStatus[Status(strings.ToUpper(v))]; !ok {
		return fmt.Errorf("parse status: invalid cage status")
	}
	return nil
}
