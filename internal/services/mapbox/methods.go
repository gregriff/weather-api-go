// helper functions to Mapbox endpoint functions
package mapbox

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/config"
)

func ForwardGeocode(mapbox http.Client, searchText string, latitude, longitude *float64) (data ForwardGeocodeResponse, err error) {
	cfg := config.Get()
	url := fmt.Sprintf(ForwardGeocodeURL, searchText, cfg.Mapbox.PublicToken)

	if latitude != nil && longitude != nil {
		url += fmt.Sprintf("&proximity=%f%%2C%f", *longitude, *latitude)
	}
	res, httpErr := mapbox.Get(url)
	if httpErr != nil {
		err = fmt.Errorf("error fetching geocoding data: %w", httpErr)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return
	}
	if jsonErr := json.NewDecoder(res.Body).Decode(&data); jsonErr != nil {
		err = fmt.Errorf("error decoding response body: %w", jsonErr)
		return
	}
	return
}
