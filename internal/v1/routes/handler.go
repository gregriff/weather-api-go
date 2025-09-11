package routes

import (
	"net/http"

	"github.com/gregriff/weather-api-go/internal/services/mapbox"
	"github.com/gregriff/weather-api-go/internal/services/nws"
)

// Handler provides the dependencies for any endpoint, and is the reciever of the endpoint handling functions
type Handler struct {
	// db        *sql.DB
	// validator *validators.RequestValidator

	// http Client that is used for all mapbox requests
	MapboxClient *http.Client

	// http Client that is used for all NWS requests
	NWSClient *http.Client
}

func NewRouteHandler() *Handler {
	return &Handler{
		// db:        db,  (*sql.DB)
		// validator: validators.New(),
		MapboxClient: mapbox.NewMapboxClient(),
		NWSClient:    nws.NewNWSClient(),
	}
}
