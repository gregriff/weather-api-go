// for static type inference of Mapbox API Response objects
package mapbox

type Geometry struct {
	Type        string    `json:"type" validate:"oneof=Point"`
	Coordinates []float64 `json:"coordinates" validate:"required"`
}

type ContextRegion struct {
	Name       string `json:"name" validate:"required"`
	RegionCode string `json:"region_code" validate:"required"`
}

type ContextPlace struct {
	Name     string `json:"name" validate:"required"`
	MapboxId string `json:"mapbox_id" validate:"required"`
}

type Context struct {
	Region ContextRegion `json:"region" validate:"required"`
	Place  ContextPlace  `json:"place" validate:"required"`
}

type FeatureCoordinates struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
}

type FeatureProperties struct {
	MapboxId       string             `json:"mapbox_id" validate:"required"`
	FeatureType    string             `json:"feature_type" validate:"required"`
	FullAddress    string             `json:"full_address" validate:"required"`
	Name           string             `json:"name" validate:"required"`
	NamePreferred  string             `json:"name_preferred" validate:"required"`
	Coordinates    FeatureCoordinates `json:"coordinates" validate:"required"`
	PlaceFormatted string             `json:"place_formatted" validate:"required"`
	Bbox           []float64          `json:"bbox" validate:"required"`
	Context        Context            `json:"context" validate:"required"`
}

type Feature struct {
	Type       string            `json:"type" validate:"oneof=Feature"`
	Id         string            `json:"id" validate:"required"`
	Geometry   Geometry          `json:"geometry" validate:"required"`
	Properties FeatureProperties `json:"properties" validate:"required"`
}

type ForwardGeocodeResponse struct {
	Type     string    `json:"type" validate:"oneof=FeatureCollection"`
	Features []Feature `json:"features" validate:"required"`
}
