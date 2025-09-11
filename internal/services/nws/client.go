// provides an http.Client for NWS requests
package nws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gregriff/weather-api-go/internal/config"
	"github.com/gregriff/weather-api-go/internal/services/shared"
)

func NewNWSClient() *http.Client {
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
