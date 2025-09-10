// for static type inference of NWS API Response objects
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
	ForecastOffice   string           `json:"forecast_office" validate:"required"`
	GridId           string           `json:"grid_id" validate:"required"`
	GridX            int              `json:"grid_x" validate:"required"`
	GridY            int              `json:"grid_y" validate:"required"`
	Forecast         string           `json:"forecast" validate:"required"`
	ForecastHourly   string           `json:"forecast_hourly" validate:"required"`
	ForecastGridData string           `json:"forecast_grid_data" validate:"required"`
	RelativeLocation RelativeLocation `json:"relative_location" validate:"required"`
}

type PointsResponse struct {
	Properties PointsProperties `json:"properties" validate:"required"`
}
