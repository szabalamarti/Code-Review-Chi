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
	// BatchCreate is a method that adds a list of vehicles to the repository
	BatchCreate(v []*Vehicle) (err error)
	// FindByColorAndYear is a method that returns a map of vehicles that match color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// Delete is a method that deletes a vehicle from the repository
	Delete(id int) (err error)
	// UpdateFuelType is a method that updates the fuel type of a vehicle in the repository
	UpdateFuelType(id int, fuelType string) (err error)
	// FindByWeightRange is a method that returns a map of vehicles that match weight range
	FindByWeightRange(minWeight, maxWeight float64) (v map[int]Vehicle, err error)
	// FindByBrandAndYearRange is a method that returns a map of vehicles that match brand and year range
	FindByBrandAndYearRange(brand string, minYear, maxYear int) (v map[int]Vehicle, err error)
}
