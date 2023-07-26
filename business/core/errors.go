package core

// Error represents custom business logic errors.
type Error string

// Error satisfies the error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrPowerDownCage represents a power down cage error.
	ErrPowerDownCage = Error("unable to power down cage whith active dinosaurs")

	// ErrInvalidCagePowerDown represents an unable to add dino to powered down cage error.
	ErrInvalidCagePowerDown = Error("unable to add dinosaurs to powered down cage")

	// ErrInvalidCageAtCapacity represents an unable to add dino to full cage error.
	ErrInvalidCageAtCapacity = Error("unable to add dinosaurs to cage at capacity")

	// ErrInvalidCageInvalidType represents an unable to add dino to diet conflict error.
	ErrInvalidCageInvalidType = Error("unable to add dinosaurs to cage with different types")

	// ErrInvalidCageInvalidSpecies represents an unable to add dino with species conflict error.
	ErrInvalidCageInvalidSpecies = Error("unable to add dinosaurs to cage with different species")

	// ErrInvalidCageInvalidRemoval represents an unable to remove dino from cage error.
	ErrInvalidCageInvalidRemoval = Error("unable to remove dinosaurs from an empty cage")

	// ErrNotFound represents an item not found.
	ErrNotFound = Error("item not found")
)
