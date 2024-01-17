package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// VehicleDefaultMock is a struct that represents the default mock for vehicles service.
type VehicleDefaultMock struct {
	mock.Mock
}

// The following methods are the implementation of the VehicleDefault interface.
func (m *VehicleDefaultMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) Create(v *internal.Vehicle) (err error) {
	args := m.Called(v)
	return args.Error(0)
}

func (m *VehicleDefaultMock) BatchCreate(v []*internal.Vehicle) (err error) {
	args := m.Called(v)
	return args.Error(0)
}

func (m *VehicleDefaultMock) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, year)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(0)
}

func (m *VehicleDefaultMock) UpdateFuelType(id int, fuelType string) (err error) {
	args := m.Called(id, fuelType)
	return args.Error(0)
}

func (m *VehicleDefaultMock) FindByWeightRange(minWeight, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	args := m.Called(minWeight, maxWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) FindByBrandAndYearRange(brand string, minYear, maxYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, minYear, maxYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
