// used to validate shape and types of NWS-related http request bodies and responses
package schemas

type Gridpoints struct {
	Office string `json:"office" validate:"required"`
	X      *int   `json:"x" validate:"required"`
	Y      *int   `json:"y" validate:"required"`

	// if user provides gridpoints when requesting a forecast, they already have this data
	City  string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
}

func (g *Gridpoints) IsEmpty() bool {
	if g.X == nil && g.Y == nil && g.Office == "" && g.City == "" && g.State == "" {
		return true
	}
	return false
}

type LocationData struct {
	Latitude   float64     `json:"latitude" validate:"required"`
	Longitude  float64     `json:"longitude" validate:"required"`
	Gridpoints *Gridpoints `json:"gridpoints"`
}

// Responses ####################################################

type PrecipitationValue struct {
	Value    any `json:"value"` // can be float, int, string, or nil?
	UnitCode any `json:"unitCode"`
}

type ForecastPeriod struct {
	Number                     int                `json:"number" validate:"required"`
	Name                       string             `json:"name" validate:"required"`
	StartTime                  string             `json:"startTime" validate:"required"`
	EndTime                    string             `json:"endTime" validate:"required"`
	IsDaytime                  *bool              `json:"isDaytime" validate:"required"`
	Temperature                *int               `json:"temperature" validate:"required"`
	TemperatureUnit            string             `json:"temperatureUnit" validate:"required"`
	TemperatureTrend           string             `json:"temperatureTrend" validate:"required"`
	ProbabilityOfPrecipitation PrecipitationValue `json:"probabilityOfPrecipitation" validate:"required"`
	WindSpeed                  string             `json:"windSpeed" validate:"required"`
	WindDirection              string             `json:"windDirection" validate:"required"`
	ShortForecast              string             `json:"shortForecast" validate:"required"`
	DetailedForecast           string             `json:"detailedForecast" validate:"required"`
	Icon                       string             `json:"icon" validate:"required"`
	IconName                   string             `json:"iconName" validate:"required"`
}

type ForecastProperties struct {
	GeneratedAt string           `json:"generatedAt" validate:"required"`
	UpdateTime  string           `json:"updateTime" validate:"required"`
	Periods     []ForecastPeriod `json:"periods" validate:"required"`
}

type ForecastResponse struct {
	Properties ForecastProperties `json:"properties" validate:"required"`
	Gridpoints Gridpoints         `json:"gridpoints" validate:"required"`
}

type HourlyForecastPrecipitationObject struct {
	Value    any    `json:"value" validate:"required"` // can be int, float, or nil
	MaxValue *int   `json:"maxValue"`
	MinValue *int   `json:"minValue"`
	UnitCode string `json:"unitCode" validate:"required"`
}

type HourlyForecastPeriod struct {
	Number                     int                               `json:"number" validate:"required"`
	Name                       string                            `json:"name" validate:"required"`
	StartTime                  string                            `json:"startTime" validate:"required"`
	EndTime                    string                            `json:"endTime" validate:"required"`
	IsDaytime                  *bool                             `json:"isDaytime" validate:"required"`
	TemperatureTrend           string                            `json:"temperatureTrend" validate:"required"`
	ProbabilityOfPrecipitation HourlyForecastPrecipitationObject `json:"probabilityOfPrecipitation" validate:"required"`
	Dewpoint                   HourlyForecastPrecipitationObject `json:"dewpoint" validate:"required"`
	RelativeHumidity           HourlyForecastPrecipitationObject `json:"relativeHumidity" validate:"required"`
	WindDirection              string                            `json:"windDirection" validate:"required"`
	ShortForecast              string                            `json:"shortForecast" validate:"required"`
	DetailedForecast           string                            `json:"detailedForecast" validate:"required"`
	Icon                       string                            `json:"icon" validate:"required"`
	IconName                   string                            `json:"iconName" validate:"required"`
}

type HourlyForecastProperties struct {
	GeneratedAt string                 `json:"generatedAt" validate:"required"`
	UpdateTime  string                 `json:"updateTime" validate:"required"`
	Periods     []HourlyForecastPeriod `json:"periods" validate:"required"`
}

type HourlyForecastResponse struct {
	Properties HourlyForecastProperties `json:"properties" validate:"required"`
}
