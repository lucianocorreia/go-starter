package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	// TemplateDir is the directory where the templates are located
	TemplateDir = "../templates"

	// TemplateExt is the extension of the templates
	TemplateExt = ".tmpl"

	// StaticDir is the directory where the static files are located
	StaticDir = "static"

	// StaticPrefix is the prefix for the static files
	StaticPrefix = "files"
)

type environment string

const (
	// EnvLocal is the local environment
	EnvLocal environment = "local"

	// EnvDev is the development environment
	EnvTest environment = "test"

	// EnvDevelop is the development environment
	EnvDevelop environment = "development"

	// EnvProduction is the production environment
	EnvProduction environment = "production"
)

// SwitchEnvironment sets the environment variable
func SwitchEnvironment(env environment) {
	if err := os.Setenv("GOSTARTER_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration
	Config struct {
		HTTP     HTTPConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
		Mail     MailConfig
	}

	// HttpConfig is the configuration for the HTTP server
	HTTPConfig struct {
		Hostname     string
		Port         uint16
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
		TLS          struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	// AppConfig stores application configuration
	AppConfig struct {
		Name          string
		Environment   environment
		EncryptionKey string
		Timeout       time.Duration
		PasswordToken struct {
			Expiration time.Duration
			Length     int
		}
		EmailVerificationTokenExpiration time.Duration
	}

	// CacheConfig stores the cache configuration
	CacheConfig struct {
		Hostname     string
		Port         uint16
		Password     string
		Database     int
		TestDatabase int
		Expiration   struct {
			StaticFile time.Duration
			Page       time.Duration
		}
	}

	// DatabaseConfig stores the database configuration
	DatabaseConfig struct {
		Hostname     string
		Port         uint16
		User         string
		Password     string
		Database     string
		TestDatabase string
	}

	// MailConfig stores the mail configuration
	MailConfig struct {
		Hostname    string
		Port        uint16
		User        string
		Password    string
		FromAddress string
	}
)

// NewCondfig returns a new Config struct
func NewConfig() (Config, error) {
	var c Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// Load env variables
	viper.SetEnvPrefix("gostarter")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil

}
