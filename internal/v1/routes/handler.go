// package routes contains the exposed API endpoints
package routes

import (
	"net/http"

	"github.com/gregriff/weather-api-go/internal/services/mapbox"
	"github.com/gregriff/weather-api-go/internal/services/nws"
)

// RouteHandler provides the dependencies for any endpoint, and is the reciever of the endpoint handling functions
type RouteHandler struct {
	// db        *sql.DB
	// validator *validators.RequestValidator

	// http Client that is used for all mapbox requests
	MapboxClient *http.Client

	// http Client that is used for all NWS requests
	NWSClient *http.Client
}

// NewRouteHandler creates the reciever for all endpoint handling functions
func NewRouteHandler() *RouteHandler {
	return &RouteHandler{
		// db:        db,  (*sql.DB)
		// validator: validators.New(),
		MapboxClient: mapbox.New(),
		NWSClient:    nws.New(),
	}
}
