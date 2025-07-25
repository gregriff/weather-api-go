// used to validate shape and types of NWS-related http request bodies and responses
package schemas

type Gridpoints struct {
	Office string `validate:"required"`
	x      int    `validate:"required"`
	y      int    `validate:"required"`

	// if user provides gridpoints when requesting a forecast, they already have this data
	City  string `validate:"required"`
	State string `validate:"required"`
}

type LocationData struct {
	Latitude   float64    `validate:"required"`
	Longitude  float64    `validate:"required"`
	Gridpoints Gridpoints `validate:"required"`
}

// Responses ####################################################

type PrecipitationValue struct {
	Value    any // can be float, int, string, or nil
	UnitCode any
}

type ForecastPeriod struct {
	Number                     int                `validate:"required"`
	Name                       string             `validate:"required"`
	StartTime                  string             `validate:"required"`
	EndTime                    string             `validate:"required"`
	IsDaytime                  bool               `validate:"required"`
	Temperature                int                `validate:"required"`
	TemperatureUnit            string             `validate:"required"`
	TemperatureTrend           string             `validate:"required"`
	ProbabilityOfPrecipitation PrecipitationValue `validate:"required"`
	WindSpeed                  string             `validate:"required"`
	WindDirection              string             `validate:"required"`
	ShortForecast              string             `validate:"required"`
	DetailedForecast           string             `validate:"required"`
	Icon                       string             `validate:"required"`
	IconName                   string             `validate:"required"`
}

type ForecastProperties struct {
	GeneratedAt string           `validate:"required"`
	UpdateTime  string           `validate:"required"`
	Periods     []ForecastPeriod `validate:"required"`
}

type ForecastResponse struct {
	Properties ForecastProperties `validate:"required"`
	Gridpoints Gridpoints         `validate:"required"`
}

type HourlyForecastPrecipitationObject struct {
	Value    any `validate:"required"` // can be int, float, or nil
	MaxValue *int
	MinValue *int
	UnitCode string `validate:"required"`
}

type HourlyForecastPeriod struct {
	Number                     int                               `validate:"required"`
	Name                       string                            `validate:"required"`
	StartTime                  string                            `validate:"required"`
	EndTime                    string                            `validate:"required"`
	IsDaytime                  bool                              `validate:"required"`
	TemperatureTrend           string                            `validate:"required"`
	ProbabilityOfPrecipitation HourlyForecastPrecipitationObject `validate:"required"`
	Dewpoint                   HourlyForecastPrecipitationObject `validate:"required"`
	RelativeHumidity           HourlyForecastPrecipitationObject `validate:"required"`
	WindDirection              string                            `validate:"required"`
	ShortForecast              string                            `validate:"required"`
	DetailedForecast           string                            `validate:"required"`
	Icon                       string                            `validate:"required"`
	IconName                   string                            `validate:"required"`
}

type HourlyForecastProperties struct {
	GeneratedAt string                 `validate:"required"`
	UpdateTime  string                 `validate:"required"`
	Periods     []HourlyForecastPeriod `validate:"required"`
}

type HourlyForecastResponse struct {
	Properties HourlyForecastProperties `validate:"required"`
}
