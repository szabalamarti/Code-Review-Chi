package repository

import "app/internal"

func ValidateVehicleMandatoryFields(v *internal.Vehicle) (err error) {
	if v.Id == 0 {
		err = internal.ErrVehicleMandatoryFields
		return
	}
	return
}
