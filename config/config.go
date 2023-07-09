package config

import (
	"fmt"
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
		Driver string
		User   string
		Pass   string
		Host   string
		Name   string
	}
)

// New returns a new Config struct
func NewConfig() Config {
	// Load .env file
	godotenv.Load()

	return Config{
		App: AppConfig{
			Name:        os.Getenv("APP_NAME"),
			Port:        os.Getenv("APP_PORT"),
			Environment: os.Getenv("APP_ENV"),
		},
		Database: DatabaseConfig{
			Driver: os.Getenv("DB_DRIVER"),
			User:   os.Getenv("DB_USER"),
			Pass:   os.Getenv("DB_PASS"),
			Host:   os.Getenv("DB_HOST"),
			Name:   os.Getenv("DB_NAME"),
		},
	}
}

// GetDBDriver returns the database driver
func (c *Config) GetDBDriver() string {
	return c.Database.Driver
}

// GetDBSource returns the database source
func (c *Config) GetDBSource() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		c.Database.User,
		c.Database.Pass,
		c.Database.Host,
		c.Database.Name,
	)
}
