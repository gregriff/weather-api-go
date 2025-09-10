// used to validate shape and types of NWS-related http request bodies and responses
package schemas

type Gridpoints struct {
	Office string `json:"office" validate:"required"`
	X      int    `json:"x" validate:"required"`
	Y      int    `json:"y" validate:"required"`

	// if user provides gridpoints when requesting a forecast, they already have this data
	City  string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
}

type LocationData struct {
	Latitude   float64    `json:"latitude" validate:"required"`
	Longitude  float64    `json:"longitude" validate:"required"`
	Gridpoints Gridpoints `json:"gridpoints" validate:"required"`
}

// Responses ####################################################

type PrecipitationValue struct {
	Value    any `json:"value"` // can be float, int, string, or nil
	UnitCode any `json:"unit_code"`
}

type ForecastPeriod struct {
	Number                     int                `json:"number" validate:"required"`
	Name                       string             `json:"name" validate:"required"`
	StartTime                  string             `json:"start_time" validate:"required"`
	EndTime                    string             `json:"end_time" validate:"required"`
	IsDaytime                  bool               `json:"is_daytime" validate:"required"`
	Temperature                int                `json:"temperature" validate:"required"`
	TemperatureUnit            string             `json:"temperature_unit" validate:"required"`
	TemperatureTrend           string             `json:"temperature_trend" validate:"required"`
	ProbabilityOfPrecipitation PrecipitationValue `json:"probability_of_precipitation" validate:"required"`
	WindSpeed                  string             `json:"wind_speed" validate:"required"`
	WindDirection              string             `json:"wind_direction" validate:"required"`
	ShortForecast              string             `json:"short_forecast" validate:"required"`
	DetailedForecast           string             `json:"detailed_forecast" validate:"required"`
	Icon                       string             `json:"icon" validate:"required"`
	IconName                   string             `json:"icon_name" validate:"required"`
}

type ForecastProperties struct {
	GeneratedAt string           `json:"generated_at" validate:"required"`
	UpdateTime  string           `json:"update_time" validate:"required"`
	Periods     []ForecastPeriod `json:"periods" validate:"required"`
}

type ForecastResponse struct {
	Properties ForecastProperties `json:"properties" validate:"required"`
	Gridpoints Gridpoints         `json:"gridpoints" validate:"required"`
}

type HourlyForecastPrecipitationObject struct {
	Value    any    `json:"value" validate:"required"` // can be int, float, or nil
	MaxValue *int   `json:"max_value"`
	MinValue *int   `json:"min_value"`
	UnitCode string `json:"unit_code" validate:"required"`
}

type HourlyForecastPeriod struct {
	Number                     int                               `json:"number" validate:"required"`
	Name                       string                            `json:"name" validate:"required"`
	StartTime                  string                            `json:"start_time" validate:"required"`
	EndTime                    string                            `json:"end_time" validate:"required"`
	IsDaytime                  bool                              `json:"is_daytime" validate:"required"`
	TemperatureTrend           string                            `json:"temperature_trend" validate:"required"`
	ProbabilityOfPrecipitation HourlyForecastPrecipitationObject `json:"probability_of_precipitation" validate:"required"`
	Dewpoint                   HourlyForecastPrecipitationObject `json:"dewpoint" validate:"required"`
	RelativeHumidity           HourlyForecastPrecipitationObject `json:"relative_humidity" validate:"required"`
	WindDirection              string                            `json:"wind_direction" validate:"required"`
	ShortForecast              string                            `json:"short_forecast" validate:"required"`
	DetailedForecast           string                            `json:"detailed_forecast" validate:"required"`
	Icon                       string                            `json:"icon" validate:"required"`
	IconName                   string                            `json:"icon_name" validate:"required"`
}

type HourlyForecastProperties struct {
	GeneratedAt string                 `json:"generated_at" validate:"required"`
	UpdateTime  string                 `json:"update_time" validate:"required"`
	Periods     []HourlyForecastPeriod `json:"periods" validate:"required"`
}

type HourlyForecastResponse struct {
	Properties HourlyForecastProperties `json:"properties" validate:"required"`
}
