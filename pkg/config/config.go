package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	App         string `json:"app"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app, ok := os.LookupEnv("APP")
	if !ok {
		panic("Unable Read Variable APP from .env File")
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("Unable Read Variable ENVIRONMENT from .env File")
	}

	version, ok := os.LookupEnv("VERSION")
	if !ok {
		panic("Unable Read Variable VERSION from .env File")
	}

	return &Config{
		App:         app,
		Environment: environment,
		Version:     version,
	}
}
