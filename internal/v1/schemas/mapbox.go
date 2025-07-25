// used to validate shape and types of Mapbox-related http request bodies and responses
package schemas

// Requests ####################################################

type GeocodeQueryData struct {
	SearchText string `validate:"required"`
	Latitude   *float64
	Longitude  *float64
}

// Responses ####################################################

type Coordinates struct {
	Longitude float64 `validate:"required"`
	Latitude  float64 `validate:"required"`
}

type PlaceData struct {
	PlaceName   string      `validate:"required"`
	RegionName  string      `validate:"required"`
	RegionCode  string      `validate:"required"`
	Coordinates Coordinates `validate:"required"`
}

type GeocodePlacesResponse struct {
	// Lowest index is highest relevance
	Results map[string]PlaceData `validate:"required"`
}
