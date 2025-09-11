package config

type APIConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type MapboxConfig struct {
	PublicToken string `json:"public_token"`
}

type NWSConfig struct {
	UserAgentID    string `json:"user_agent_identifier"`
	UserAgentEmail string `json:"user_agent_email"`

	TestLat  float64 `json:"test_lat"`
	TestLong float64 `json:"test_long"`
}

type Config struct {
	Debug bool `json:"debug"`

	Api    APIConfig    `json:"api"`
	Mapbox MapboxConfig `json:"mapbox"`
	NWS    NWSConfig    `json:"nws"`
}
