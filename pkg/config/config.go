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
	err := godotenv.Load(".config.env")
	if err != nil {
		panic("Error loading .config.env file")
	}

	app, ok := os.LookupEnv("APP")
	if !ok {
		panic("Unable Read Variable APP from .config.env File")
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("Unable Read Variable ENVIRONMENT from .config.env File")
	}

	version, ok := os.LookupEnv("VERSION")
	if !ok {
		panic("Unable Read Variable VERSION from .config.env File")
	}

	return &Config{
		App:         app,
		Environment: environment,
		Version:     version,
	}
}
