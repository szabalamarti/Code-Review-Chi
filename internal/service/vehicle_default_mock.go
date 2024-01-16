package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type VehicleDefaultMock struct {
	mock.Mock
}

func (m *VehicleDefaultMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) Create(v *internal.Vehicle) (err error) {
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
