// helper functions to NWS endpoint functions
package nws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/v1/schemas"
)

// fetchGridpoints expects string lat and long returned by FormatCoordinates
func fetchGridpoints(nws *http.Client, latitude, longitude string) (data PointsResponse, err error) {
	var res *http.Response

	url := fmt.Sprintf(PointsURL, latitude, longitude)
	res, err = nws.Get(url)
	if err != nil {
		err = fmt.Errorf("error during request: %w", err)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if statusCode := res.StatusCode; statusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %d", statusCode)
		return
	}

	if jsonErr := json.NewDecoder(res.Body).Decode(&data); jsonErr != nil {
		err = fmt.Errorf("error decoding response body: %w", jsonErr)
		return
	}
	return
}

// GetGridpoints expects string lat and long returned by FormatCoordinates
func GetGridpoints(nws *http.Client, latitude, longitude string) (schemas.Gridpoints, error) {
	var data = schemas.Gridpoints{}

	res, err := fetchGridpoints(nws, latitude, longitude)
	if err != nil {
		err = fmt.Errorf("error fetching gridpoints: %w", err)
		return data, err
	}
	gridpointProps := res.Properties
	locationProps := gridpointProps.RelativeLocation.Properties

	data.Office = gridpointProps.Cwa
	data.X = &gridpointProps.GridX
	data.Y = &gridpointProps.GridY
	data.City = locationProps.City
	data.State = locationProps.State
	return data, nil
}

// GetForecast expects string lat and long returned by FormatCoordinates
func GetForecast(nws *http.Client, latitude, longitude string, gridpoints schemas.Gridpoints) (data schemas.ForecastResponse, err error) {
	if gridpoints.IsEmpty() {
		gridpoints, err = GetGridpoints(nws, latitude, longitude)
		if err != nil {
			return
		}
	}

	urlParams := fmt.Sprintf(GridpointURLParams, gridpoints.Office, *gridpoints.X, *gridpoints.Y)
	url := fmt.Sprintf(ForecastURL, urlParams)
	res, httpErr := nws.Get(url)
	if httpErr != nil {
		err = fmt.Errorf("error fetching forecast: %w", httpErr)
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

	// prepare response
	data.Gridpoints = gridpoints
	if newIconNames, ok := SetIconNames(data.Properties.Periods).([]schemas.ForecastPeriod); ok {
		data.Properties.Periods = newIconNames
	} else {
		panic("logic error with typing in SetIconNames")
	}
	return
}

func GetHourlyForecast(nws *http.Client, latitude, longitude string, gridpoints schemas.Gridpoints) (data schemas.HourlyForecastResponse, err error) {
	if gridpoints.IsEmpty() {
		gridpoints, err = GetGridpoints(nws, latitude, longitude)
		if err != nil {
			return
		}
	}

	urlParams := fmt.Sprintf(GridpointURLParams, gridpoints.Office, *gridpoints.X, *gridpoints.Y)
	url := fmt.Sprintf(HourlyForecastURL, urlParams)
	res, httpErr := nws.Get(url)
	if httpErr != nil {
		err = fmt.Errorf("error fetching hourly forecast: %w", httpErr)
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

	// prepare response
	if newIconNames, ok := SetIconNames(data.Properties.Periods).([]schemas.HourlyForecastPeriod); ok {
		data.Properties.Periods = newIconNames
	} else {
		panic("logic error with typing in SetIconNames")
	}
	return
}
