package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/services/mapbox"
	"github.com/gregriff/weather-api-go/internal/v1/schemas"
)

func (h *RouteHandler) GeocodePlace(w http.ResponseWriter, r *http.Request) {
	query := schemas.GeocodeQueryData{}
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mapboxRes, err := mapbox.ForwardGeocode(*h.MapboxClient, query.SearchText, query.Latitude, query.Longitude)
	if err != nil {
		log.Println(fmt.Errorf("GeocodePlace Error: %w", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var (
		coords  schemas.Coordinates
		context mapbox.Context
		place   mapbox.ContextPlace
		region  mapbox.ContextRegion
	)
	results := make(map[string]schemas.PlaceData, 5)

	for _, feature := range mapboxRes.Features {
		coords = schemas.Coordinates{ // TODO: ensure indexing is correct
			Longitude: feature.Geometry.Coordinates[0],
			Latitude:  feature.Geometry.Coordinates[1],
		}
		context = feature.Properties.Context
		place = context.Place
		region = context.Region

		results[place.MapboxId] = schemas.PlaceData{
			Coordinates: coords,
			PlaceName:   place.Name,
			RegionName:  region.Name,
			RegionCode:  region.RegionCode,
		}
	}

	data := schemas.GeocodePlacesResponse{Results: results}
	WriteJSON(w, data, 200)
}
