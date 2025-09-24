package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

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
	var configFile string
	if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
		configFile = configPath
	} else {
		configFile = "dev.env.json"
	}

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file %s: %v", configFile, err)
	}
	defer func() {
		_ = file.Close()
	}()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file %s: %v", configFile, err)
	}

	return &config
}
