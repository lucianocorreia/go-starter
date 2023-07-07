package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Config is a struct to hold the configuration for the application
	Config struct {
		App      AppConfig
		Database DatabaseConfig
	}

	// AppConfig is a struct to hold the configuration for the application
	AppConfig struct {
		Name string
		Port string

		Environment string
	}

	// DatabaseConfig is a struct to hold the configuration for the database
	DatabaseConfig struct {
		User string
		Pass string
		Host string
		Name string
	}
)

// New returns a new Config struct
func New() *Config {
	// Load .env file
	godotenv.Load()

	return &Config{
		App: AppConfig{
			Name:        os.Getenv("APP_NAME"),
			Port:        os.Getenv("APP_PORT"),
			Environment: os.Getenv("APP_ENV"),
		},
	}
}
