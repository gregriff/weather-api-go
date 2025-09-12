// helper functions to NWS endpoint functions
package nws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/v1/schemas"
)

// GetGridpointsRaw expects string lat and long returned by FormatCoordinates
func GetGridpointsRaw(nws *http.Client, latitude, longitude string) (data PointsResponse, httpErr error) {
	url := fmt.Sprintf(PointsURL, latitude, longitude)
	res, err := nws.Get(url)
	if err != nil {
		log.Printf("res: %#v", res)
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}
	defer res.Body.Close()

	if statusCode := res.StatusCode; statusCode != http.StatusOK {
		if statusCode == 400 {
			httpErr = errors.New("GET /gridpoints failed with 400")
			log.Printf("ERROR: %v", httpErr)
			return
		}
		httpErr = errors.New("GET /gridpoints returned non-200 status")
		log.Printf("ERROR: %v", httpErr)
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		httpErr = errors.New("Json parse failed")
		log.Printf("ERROR: %v", httpErr)
		return
	}
	return
}

// GetGridpoints expects string lat and long returned by FormatCoordinates
func GetGridpoints(nws *http.Client, latitude, longitude string) (data *schemas.Gridpoints, httpErr error) {
	res, err := GetGridpointsRaw(nws, latitude, longitude)
	if err != nil {
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}
	gridpointProps := res.Properties
	locationProps := gridpointProps.RelativeLocation.Properties

	// TODO: ensure compiler is dereferencing properly here
	data = &schemas.Gridpoints{}
	data.Office = gridpointProps.Cwa
	data.X = gridpointProps.GridX
	data.Y = gridpointProps.GridY
	data.City = locationProps.City
	data.State = locationProps.State
	return
}

// GetForecastRaw expects string lat and long returned by FormatCoordinates
func GetForecastRaw(nws *http.Client, latitude, longitude string, gridpoints *schemas.Gridpoints) (data schemas.ForecastResponse, httpErr error) {
	if gridpoints == nil || gridpoints.IsEmpty() {
		gridpoints, httpErr = GetGridpoints(nws, latitude, longitude)
		if httpErr != nil {
			log.Printf("ERROR: %v", httpErr)
			return
		}
	}

	urlParams := fmt.Sprintf(GridpointURLParams, gridpoints.Office, gridpoints.X, gridpoints.Y)
	url := fmt.Sprintf(ForecastURL, urlParams)
	res, err := nws.Get(url)
	if err != nil {
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		httpErr = fmt.Errorf("Request returned non-200 status: %s", res.Status)
		log.Printf("ERROR: %v", httpErr)
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		httpErr = err
		log.Printf("ERROR: %v", httpErr)
		return
	}

	// prepare response
	data.Gridpoints = *gridpoints
	if newIconNames, ok := SetIconNames(data.Properties.Periods).([]schemas.ForecastPeriod); ok {
		data.Properties.Periods = newIconNames
	} else {
		panic("logic error with typing in SetIconNames")
	}
	return
}

func GetHourlyForecastRaw(nws *http.Client, latitude, longitude string, gridpoints *schemas.Gridpoints) (data schemas.HourlyForecastResponse, httpErr error) {
	if gridpoints == nil || gridpoints.IsEmpty() {
		gridpoints, httpErr = GetGridpoints(nws, latitude, longitude)
		if httpErr != nil {
			log.Printf("ERROR: %v", httpErr)
			return
		}
	}

	urlParams := fmt.Sprintf(GridpointURLParams, gridpoints.Office, gridpoints.X, gridpoints.Y)
	url := fmt.Sprintf(HourlyForecastURL, urlParams)
	res, err := nws.Get(url)
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

	// prepare response
	if newIconNames, ok := SetIconNames(data.Properties.Periods).([]schemas.HourlyForecastPeriod); ok {
		data.Properties.Periods = newIconNames
	} else {
		panic("logic error with typing in SetIconNames")
	}
	return
}
