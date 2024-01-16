package internal

import "errors"

var (
	// ErrInternalServer is an error that represents an internal server error
	ErrInternalServer = errors.New("internal server error")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that adds a vehicle to the repository
	Create(v *Vehicle) (err error)
}
