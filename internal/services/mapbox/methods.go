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
		fmt.Println("FORWARD GEOCODE: lat: ", *latitude, "long: ", *longitude)
		url += fmt.Sprintf("&proximity=%d%%2C%d", longitude, latitude)
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
