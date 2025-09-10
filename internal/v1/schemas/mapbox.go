// used to validate shape and types of Mapbox-related http request bodies and responses
package schemas

// Requests ####################################################

type GeocodeQueryData struct {
	SearchText string   `json:"search_text" validate:"required"`
	Latitude   *float64 `json:"latitude"`
	Longitude  *float64 `json:"longitude"`
}

// Responses ####################################################

type Coordinates struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
}

type PlaceData struct {
	PlaceName   string      `json:"place_name" validate:"required"`
	RegionName  string      `json:"region_name" validate:"required"`
	RegionCode  string      `json:"region_code" validate:"required"`
	Coordinates Coordinates `json:"coordinates" validate:"required"`
}

type GeocodePlacesResponse struct {
	// Lowest index is highest relevance
	Results map[string]PlaceData `json:"results" validate:"required"`
}
