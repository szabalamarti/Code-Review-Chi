package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that adds a vehicle to the repository
func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	if err != nil {
		switch err {
		case internal.ErrVehicleAlreadyExists:
			return
		case internal.ErrVehicleMandatoryFields:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// BatchCreate is a method that adds a list of vehicles to the repository
func (s *VehicleDefault) BatchCreate(v []*internal.Vehicle) (err error) {
	err = s.rp.BatchCreate(v)
	if err != nil {
		switch err {
		case internal.ErrVehicleAlreadyExists:
			return
		case internal.ErrVehicleMandatoryFields:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// FindByColorAndYear is a method that returns a map of vehicles that match color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// Delete is a method that deletes a vehicle from the repository
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrVehicleNotFound:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// UpdateFuelType is a method that updates the fuel type of a vehicle in the repository
func (s *VehicleDefault) UpdateFuelType(id int, fuelType string) (err error) {
	err = s.rp.UpdateFuelType(id, fuelType)
	if err != nil {
		switch err {
		case internal.ErrVehicleNotFound:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// FindByWeightRange is a method that returns a map of vehicles that match weight range
func (s *VehicleDefault) FindByWeightRange(minWeight, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByWeightRange(minWeight, maxWeight)
	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match brand and year range
func (s *VehicleDefault) FindByBrandAndYearRange(brand string, minYear, maxYear int) (v map[int]internal.Vehicle, err error) {
	// call repository method
	v, err = s.rp.FindByBrandAndYearRange(brand, minYear, maxYear)
	// handle errors
	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			return
		default:
			err = internal.ErrInternalServer
		}
	}
	return
}
