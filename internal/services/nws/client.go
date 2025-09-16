// package nws provides functionality to communicate with the NWS API: https://api.weather.gov/
package nws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gregriff/weather-api-go/internal/config"
	"github.com/gregriff/weather-api-go/internal/services/shared"
)

// New creates an http.Client for NWS requests
func New() *http.Client {
	cfg := config.Get()
	nwsTransport := shared.Transport{
		BaseURL: BaseURL,
		Headers: map[string]string{
			"User-Agent": fmt.Sprintf("%s, %s", cfg.NWS.UserAgentID, cfg.NWS.UserAgentEmail),
		},
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: &nwsTransport,
	}
}
