package internal

import "errors"

var (
	// ErrVehicleAlreadyExists is an error that represents that the vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	// ErrVehicleMandatoryFields is an error that represents that the vehicle is missing mandatory fields
	ErrVehicleMandatoryFields = errors.New("vehicle missing mandatory fields")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that adds a vehicle to the repository
	Create(v *Vehicle) (err error)
}
