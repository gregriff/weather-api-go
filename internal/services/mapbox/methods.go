// helper functions to Mapbox endpoint functions
package mapbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/config"
)

func ForwardGeocode(mapbox http.Client, searchText string, latitude, longitude *float64) (data ForwardGeocodeResponse, httpErr error) {
	cfg := config.Get()
	url := fmt.Sprintf(ForwardGeocodeURL, searchText, cfg.Mapbox.PublicToken)

	if latitude != nil && longitude != nil {
		url += fmt.Sprintf("&proximity=%f%%2C%f", *longitude, *latitude)
	}
	res, err := mapbox.Get(url)
	if err != nil {
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		httpErr = errors.New("Request returned non-200 status")
		log.Printf("ERROR: %v", httpErr)
		return
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}
	return
}
