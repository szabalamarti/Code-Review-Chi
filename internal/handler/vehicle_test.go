package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetByColorAndYear(t *testing.T) {
	t.Run("should return status code 200 when vehicles are found", func(t *testing.T) {
		// ARRANGE
		// - vehicles
		vehicles := map[int]internal.Vehicle{
			1001: *internal.NewVehicle(1001, "Toyota", "Corolla", "ABC-1234", "Blue", 2020, 5, 180.0, "Gasoline", "Automatic", 1300.0, 1.45, 4.62, 1.77),
			1002: *internal.NewVehicle(1002, "Ford", "Fiesta", "DEF-5678", "Red", 2019, 5, 180.0, "Gasoline", "Automatic", 1300.0, 1.45, 4.62, 1.77),
		}
		// - expected response with message and data
		expectedResponse := `{"message":"Vehículos encontrados exitosamente.","data":{"1001":{"id":1001,"brand":"Toyota","model":"Corolla","registration":"ABC-1234","color":"Blue","year":2020,"passengers":5,"max_speed":180,"fuel_type":"Gasoline","transmission":"Automatic","weight":1300,"height":1.45,"length":4.62,"width":1.77},"1002":{"id":1002,"brand":"Ford","model":"Fiesta","registration":"DEF-5678","color":"Red","year":2019,"passengers":5,"max_speed":180,"fuel_type":"Gasoline","transmission":"Automatic","weight":1300,"height":1.45,"length":4.62,"width":1.77}}}`

		// - service mock
		service := new(service.VehicleDefaultMock)
		// define mock behavior
		service.On("FindByColorAndYear", "Blue", 2020).Return(vehicles, nil)

		// - Request
		req := httptest.NewRequest(http.MethodGet, "/vehicles?color=Blue&year=2020", nil)
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("color", "Blue")
		routeContext.URLParams.Add("year", "2020")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		// - ResponseRecorder
		rr := httptest.NewRecorder()
		// - Handler
		h := handler.NewVehicleDefault(service)
		// - HandlerFunc
		reqHandler := http.HandlerFunc(h.GetByColorAndYear())

		// ACT
		reqHandler.ServeHTTP(rr, req)

		// ASSERT
		// - status code
		require.Equal(t, http.StatusOK, rr.Code)
		require.JSONEq(t, expectedResponse, strings.TrimSpace(rr.Body.String()))
	})
}

func TestCreate(t *testing.T) {
	t.Run("should return status code 201 when vehicle is created", func(t *testing.T) {
		// ARRANGE
		// - vehicle
		vehicle := internal.NewVehicle(1001, "Toyota", "Corolla", "ABC-1234", "Blue", 2020, 5, 180.0, "Gasoline", "Automatic", 1300.0, 1.45, 4.62, 1.77)

		// - service mock
		service := new(service.VehicleDefaultMock)
		// define mock behavior
		service.On("Create", mock.MatchedBy(func(v *internal.Vehicle) bool {
			return reflect.DeepEqual(v, vehicle)
		})).Return(nil)

		// - Request Body
		reqBody := `{
			"id": 1001,
			"brand": "Toyota",
			"model": "Corolla",
			"registration": "ABC-1234",
			"color": "Blue",
			"year": 2020,
			"passengers": 5,
			"max_speed": 180.0,
			"fuel_type": "Gasoline",
			"transmission": "Automatic",
			"weight": 1300.0,
			"height": 1.45,
			"length": 4.62,
			"width": 1.77
		}`

		// - handler
		h := handler.NewVehicleDefault(service)
		reqHandler := http.HandlerFunc(h.Create())

		// - request
		req := httptest.NewRequest(http.MethodPost, "/vehicles", strings.NewReader(reqBody))
		// - response recorder
		w := httptest.NewRecorder()

		// ACT
		reqHandler.ServeHTTP(w, req)

		// ASSERT
		// - status code
		require.Equal(t, http.StatusCreated, w.Code)
	})

}
