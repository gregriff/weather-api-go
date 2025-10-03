// helper functions for all endpoint handling functions
package routes

import (
	"net/http"

	"github.com/gregriff/weather-api-go/internal/validation"
)

// WriteValidJSON validates the data, writing an error to the response if encountered.
// Otherwise it writes the data as JSON to the response and the status
func WriteValidJSON(w http.ResponseWriter, data any, status int) {
	if err := validation.ValidateAndEncode(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
