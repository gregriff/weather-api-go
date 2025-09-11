// provides an http.Client for Mapbox requests
package mapbox

import (
	"net/http"
	"time"

	"github.com/gregriff/weather-api-go/internal/services/shared"
)

func NewMapboxClient() *http.Client {
	mapboxTransport := shared.Transport{
		BaseURL: BaseURL,
	}

	return &http.Client{
		Timeout:   5 * time.Second,
		Transport: &mapboxTransport,
	}
}
