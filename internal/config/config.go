package config

import "os"

type Config struct {
	Port string
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
