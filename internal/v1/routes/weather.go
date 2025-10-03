package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/services/nws"
	"github.com/gregriff/weather-api-go/internal/v1/schemas"
	"github.com/gregriff/weather-api-go/internal/validation"
)

func (h *RouteHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
	query := schemas.LocationData{}
	if err := validation.DecodeAndValidate(r.Body, &query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lat, long := nws.FormatCoordinates(query.Latitude, query.Longitude)
	res, err := nws.GetForecast(h.NWSClient, lat, long, query.Gridpoints)
	if err != nil {
		log.Println(fmt.Errorf("GetForecast Error: %w", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteValidJSON(w, &res, 200)
}

func (h *RouteHandler) GetHourlyForecast(w http.ResponseWriter, r *http.Request) {
	query := schemas.LocationData{}
	if err := validation.DecodeAndValidate(r.Body, &query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lat, long := nws.FormatCoordinates(query.Latitude, query.Longitude)
	res, err := nws.GetHourlyForecast(h.NWSClient, lat, long, query.Gridpoints)
	if err != nil {
		log.Println(fmt.Errorf("GetHourlyForecast Error: %w", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteValidJSON(w, &res, 200)
}
