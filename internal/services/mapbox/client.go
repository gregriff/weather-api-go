// package mapbox provides functionality to communicate with the Mapbox Geocoding API: https://api.mapbox.com/search/geocode/v6
package mapbox

import (
	"net/http"
	"time"

	"github.com/gregriff/weather-api-go/internal/services/shared"
)

// New provides an http.Client for Mapbox requests
func New() *http.Client {
	mapboxTransport := shared.Transport{
		BaseURL:               BaseURL,
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}

	return &http.Client{
		Timeout:   5 * time.Second,
		Transport: &mapboxTransport,
	}
}
