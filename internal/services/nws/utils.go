// more lower-level (not operating on requests or responses) helpers to NWS endpoint functions
package nws

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregriff/weather-api-go/internal/v1/schemas"
)

// To comply with NWS API, we need to ensure coordinates have a maximum of 4 digits after the decimal, but no trailing zeroes
func FormatCoordinates(latitude, longitude any) (lat, long string) {
	lat, long = formatCoordinate(latitude), formatCoordinate(longitude)
	return
}

func formatCoordinate(coord any) (formatted string) {
	var f float64

	switch v := coord.(type) {
	case string:
		var err error
		f, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return "0" // or handle error as appropriate
		}
	case int:
		f = float64(v)
	case float64:
		f = v
	case float32:
		f = float64(v)
	default:
		return "0" // or handle error as appropriate
	}

	formatted = fmt.Sprintf("%.4f", f)            // Format to 4 decimal places
	formatted = strings.TrimRight(formatted, "0") // Remove trailing zeros
	formatted = strings.TrimRight(formatted, ".") // Remove trailing decimal point if all decimals were zeros

	return
}

// SetIconNames returns a new periods struct with correct IconName mappings to what the frontend can render
func SetIconNames(periods any) any {
	switch periods := periods.(type) {
	case []schemas.ForecastPeriod:
		for idx, period := range periods {
			periods[idx].IconName = getIconName(period.Icon)
		}
	case []schemas.HourlyForecastPeriod:
		for idx, period := range periods {
			periods[idx].IconName = getIconName(period.Icon)
		}
	}
	return periods
}

// getIconName turns a NWS icon string into one that the frontend looks for when it decides what to render
func getIconName(icon string) (newIconName string) {
	iconPrefix := strings.Split(icon, "/")[1]
	newIconName = strings.Split(iconPrefix, ",")[0]

	if strings.Contains(newIconName, "?") {
		newIconName = strings.Split(newIconName, "?")[0]
	}
	return
}
