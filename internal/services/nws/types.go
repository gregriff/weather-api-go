// for static type inference of NWS API Response objects
package nws

type RelativeLocationProperties struct {
	City  string `validate:"required"`
	State string `validate:"required"`
}

type RelativeLocation struct {
	Properties RelativeLocationProperties `validate:"required"`
}

type PointsProperties struct {
	Cwa              string           `validate:"required"`
	ForecastOffice   string           `validate:"required"`
	GridId           string           `validate:"required"`
	GridX            int              `validate:"required"`
	GridY            int              `validate:"required"`
	Forecast         string           `validate:"required"`
	ForecastHourly   string           `validate:"required"`
	ForecastGridData string           `validate:"required"`
	RelativeLocation RelativeLocation `validate:"required"`
}

type PointsResponse struct {
	Properties PointsProperties `validate:"required"`
}
