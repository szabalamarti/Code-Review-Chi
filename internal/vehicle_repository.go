package internal

import "errors"

var (
	// ErrVehicleAlreadyExists is an error that represents that the vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	// ErrVehicleMandatoryFields is an error that represents that the vehicle is missing mandatory fields
	ErrVehicleMandatoryFields = errors.New("vehicle missing mandatory fields")
	// ErrVehiclesNotFound is an error that represents that no vehicles were found with the given criteria
	ErrVehiclesNotFound = errors.New("vehicles not found")
	// ErrVehicleNotFound is an error that represents that the vehicle was not found
	ErrVehicleNotFound = errors.New("vehicle not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that adds a vehicle to the repository
	Create(v *Vehicle) (err error)
	// FindByColorAndYear is a method that returns a map of vehicles that match color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// Delete is a method that deletes a vehicle from the repository
	Delete(id int) (err error)
	// UpdateFuelType is a method that updates the fuel type of a vehicle in the repository
	UpdateFuelType(id int, fuelType string) (err error)
	// FindByWeightRange is a method that returns a map of vehicles that match weight range
	FindByWeightRange(minWeight, maxWeight float64) (v map[int]Vehicle, err error)
}
