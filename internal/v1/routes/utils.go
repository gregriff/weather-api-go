// helper functions for all endpoint handling functions
package routes

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes JSON to a response, and a status code to the header
func WriteJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
