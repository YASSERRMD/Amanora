package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment string
	HTTPAddr    string
	LogLevel    string
}

func Load() Config {
	return Config{
		Environment: getEnv("AMANORA_ENV", "development"),
		HTTPAddr:    getEnv("AMANORA_HTTP_ADDR", ":8080"),
		LogLevel:    getEnv("AMANORA_LOG_LEVEL", "info"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}
