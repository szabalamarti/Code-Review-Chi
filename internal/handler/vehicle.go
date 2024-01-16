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
