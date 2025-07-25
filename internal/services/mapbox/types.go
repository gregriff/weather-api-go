// for static type inference of Mapbox API Response objects
package mapbox

type Geometry struct {
	Type        string    `validate:"oneof=Point"`
	Coordinates []float64 `validate:"required"`
}

type ContextRegion struct {
	Name       string `validate:"required"`
	RegionCode string `validate:"required"`
}

type ContextPlace struct {
	Name     string `validate:"required"`
	MapboxId string `validate:"required"`
}

type Context struct {
	Region ContextRegion `validate:"required"`
	Place  ContextPlace  `validate:"required"`
}

type FeatureCoordinates struct {
	Longitude float64 `validate:"required"`
	Latitude  float64 `validate:"required"`
}

type FeatureProperties struct {
	MapboxId       string             `validate:"required"`
	FeatureType    string             `validate:"required"`
	FullAddress    string             `validate:"required"`
	Name           string             `validate:"required"`
	NamePreferred  string             `validate:"required"`
	Coordinates    FeatureCoordinates `validate:"required"`
	PlaceFormatted string             `validate:"required"`
	Bbox           []float64          `validate:"required"`
	Context        Context            `validate:"required"`
}

type Feature struct {
	Type       string            `validate:"oneof=Feature"`
	Id         string            `validate:"required"`
	Geometry   Geometry          `validate:"required"`
	Properties FeatureProperties `validate:"required"`
}

type ForwardGeocodeResponse struct {
	Type     string    `validate:"oneof=FeatureCollection"`
	Features []Feature `validate:"required"`
}
