// NWS API Response types
package nws

type RelativeLocationProperties struct {
	City  string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
}

type RelativeLocation struct {
	Properties RelativeLocationProperties `json:"properties" validate:"required"`
}

type PointsProperties struct {
	Cwa              string           `json:"cwa" validate:"required"`
	ForecastOffice   string           `json:"forecastOffice" validate:"required"`
	GridId           string           `json:"gridId" validate:"required"`
	GridX            int              `json:"gridX" validate:"required"`
	GridY            int              `json:"gridY" validate:"required"`
	Forecast         string           `json:"forecast" validate:"required"`
	ForecastHourly   string           `json:"forecastHourly" validate:"required"`
	ForecastGridData string           `json:"forecastGridData" validate:"required"`
	RelativeLocation RelativeLocation `json:"relativeLocation" validate:"required"`
}

type PointsResponse struct {
	Properties PointsProperties `json:"properties" validate:"required"`
}
