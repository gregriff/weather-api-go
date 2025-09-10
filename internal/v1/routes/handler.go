package routes

import (
	"net/http"
	"time"
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

func New() *Handler {
	return &Handler{
		// db:        db,  (*sql.DB)
		// validator: validators.New(),
		MapboxClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NWSClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
