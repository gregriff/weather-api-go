package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type MapboxConfig struct {
	PublicToken string `json:"public_token"`
}

type NWSConfig struct {
	TestLat  int `json:"test_lat"`
	TestLong int `json:"test_long"`
}

type Config struct {
	Port string `json:"port"`

	Mapbox MapboxConfig `json:"mapbox"`
	NWS    NWSConfig    `json:"nws"`
}

var (
	instance *Config
	once     sync.Once
)

// Get returns the singleton config instance
func Get() *Config {
	once.Do(func() {
		instance = load()
	})
	return instance
}

// load reads and parses the config file
func load() *Config {
	configFile := getConfigFilePath("dev.env.json")

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file %s: %v", configFile, err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file %s: %v", configFile, err)
	}

	return &config
}

// getConfigFilePath determines the config file location
func getConfigFilePath(defaultPath string) string {
	if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
		return configPath
	}
	return defaultPath
}
