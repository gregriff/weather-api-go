package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/config"
	"github.com/gregriff/weather-api-go/internal/services/nws"
	"github.com/gregriff/weather-api-go/internal/v1/schemas"
)

func (h *RouteHandler) TestForecast(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	lat, long := nws.FormatCoordinates(cfg.NWS.TestLat, cfg.NWS.TestLong)
	res, err := nws.GetForecastRaw(h.NWSClient, lat, long, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, res, 200)
}

func (h *RouteHandler) TestGridpoints(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	lat, long := nws.FormatCoordinates(cfg.NWS.TestLat, cfg.NWS.TestLong)
	res, err := nws.GetGridpointsRaw(h.NWSClient, lat, long)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, res, 200)
}

func (h *RouteHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
	query := schemas.LocationData{}
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lat, long := nws.FormatCoordinates(query.Latitude, query.Longitude)
	res, err := nws.GetForecastRaw(h.NWSClient, lat, long, &query.Gridpoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, res, 200)
}

func (h *RouteHandler) GetHourlyForecast(w http.ResponseWriter, r *http.Request) {
	query := schemas.LocationData{}
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lat, long := nws.FormatCoordinates(query.Latitude, query.Longitude)
	res, err := nws.GetHourlyForecastRaw(h.NWSClient, lat, long, &query.Gridpoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, res, 200)
}
