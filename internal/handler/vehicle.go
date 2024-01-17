package handler

import (
	"app/internal"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Datos del vehículo mal formados o incompletos.")
			return
		}

		// process
		vehicle := internal.NewVehicle(
			reqBody.ID,
			reqBody.Brand,
			reqBody.Model,
			reqBody.Registration,
			reqBody.Color,
			reqBody.FabricationYear,
			reqBody.Capacity,
			reqBody.MaxSpeed,
			reqBody.FuelType,
			reqBody.Transmission,
			reqBody.Weight,
			reqBody.Height,
			reqBody.Length,
			reqBody.Width,
		)

		err = h.sv.Create(vehicle)
		if err != nil {
			switch err {
			case internal.ErrVehicleAlreadyExists:
				response.Error(w, http.StatusConflict, "Identificador del vehículo ya existente.")
				return
			case internal.ErrVehicleMandatoryFields:
				response.Error(w, http.StatusBadRequest, "Datos del vehículo mal formados o incompletos.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehículo creado exitosamente.",
		})
	}
}

// VehicleBatchJSON is a struct that represents a list of vehicles in JSON format
type VehicleBatchJSON struct {
	Vehicles []VehicleJSON `json:"vehicles"`
}

// BatchCreate is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) BatchCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody VehicleBatchJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Datos de algún vehículo mal formados o incompletos.")
			return
		}

		// process
		// make a slice of pointers
		vehicles := make([]*internal.Vehicle, len(reqBody.Vehicles))
		for i, v := range reqBody.Vehicles {
			vehicles[i] = internal.NewVehicle(
				v.ID,
				v.Brand,
				v.Model,
				v.Registration,
				v.Color,
				v.FabricationYear,
				v.Capacity,
				v.MaxSpeed,
				v.FuelType,
				v.Transmission,
				v.Weight,
				v.Height,
				v.Length,
				v.Width,
			)
		}

		// call the service to create the vehicles
		err = h.sv.BatchCreate(vehicles)
		if err != nil {
			switch err {
			case internal.ErrVehicleAlreadyExists:
				response.Error(w, http.StatusConflict, "Algún vehículo tiene un identificador ya existente.")
				return
			case internal.ErrVehicleMandatoryFields:
				response.Error(w, http.StatusBadRequest, "Datos de algún vehículo mal formados o incompletos.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehículos creados exitosamente.",
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get color and year from URL using chi
		color := chi.URLParam(r, "color")
		yearString := chi.URLParam(r, "year")
		year, err := strconv.Atoi(yearString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Año mal formado.")
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			switch err {
			case internal.ErrVehiclesNotFound:
				response.Error(w, http.StatusNotFound, "No se encontraron vehículos con esos criterios.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Vehículos encontrados exitosamente.",
			"data":    data,
		})
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/{id}
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from URL using chi
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Identificador mal formado.")
			return
		}

		// process
		// - delete vehicle by id
		err = h.sv.Delete(id)
		if err != nil {
			switch err {
			case internal.ErrVehicleNotFound:
				response.Error(w, http.StatusNotFound, "No se encontró el vehículo con ese identificador.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		response.Text(w, http.StatusNoContent, "Vehículo eliminado exitosamente.")
	}
}

// UpdateFuelTypeJSON is a struct that represents the request body for the route PUT /vehicles/{id}/fuel_type
type UpdateFuelTypeJSON struct {
	FuelType string `json:"fuel_type"`
}

// UpdateFuelType is a method that returns a handler for the route PUT /vehicles/{id}/fuel_type
func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from URL using chi
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Identificador mal formado.")
			return
		}

		// - get fuel type from request body
		var reqBody UpdateFuelTypeJSON
		err = json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Tipo de combustible mal formado o no admitido.")
			return
		}

		// process
		// - update fuel type by id
		err = h.sv.UpdateFuelType(id, reqBody.FuelType)
		if err != nil {
			switch err {
			case internal.ErrVehicleNotFound:
				response.Error(w, http.StatusNotFound, "No se encontró el vehículo con ese identificador.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Tipo de combustible del vehículo actualizado exitosamente.",
		})
	}
}

// GetByWeightRange is a method that returns a handler for the route GET /vehicles/weight?min={weight_min}&max={weight_max}
func (h *VehicleDefault) GetByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get min and max from URL using chi
		minString := r.URL.Query().Get("min")
		maxString := r.URL.Query().Get("max")

		// convert them to float64
		min, err := strconv.ParseFloat(minString, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Peso mínimo mal formado.")
			return
		}
		max, err := strconv.ParseFloat(maxString, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Peso máximo mal formado.")
			return
		}

		// process
		// - get vehicles by weight range
		v, err := h.sv.FindByWeightRange(min, max)
		if err != nil {
			switch err {
			case internal.ErrVehiclesNotFound:
				response.Error(w, http.StatusNotFound, "No se encontraron vehículos con esos criterios.")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "Algo ha salido mal.")
				return
			}
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Vehículos encontrados exitosamente.",
			"data":    data,
		})
	}
}
