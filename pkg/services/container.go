package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lucianocorreia/go-starter/config"
)

// Container is the container for the services
type Container struct {
	// Validator stores a validator
	Validator *Validator

	// Web is the web server
	Web *echo.Echo

	// Config is the configuration
	Config *config.Config

	// TemplateRenderer stores a service to easily render and cache templates
	TemplateRenderer *TemplateRenderer
}

// NewContainer creates a new container
func NewContainer() *Container {
	c := new(Container)

	c.initConfig()
	c.initValidator()

	c.initWeb()

	c.initTemplateRenderer()

	return c
}

// Shutdown shuts down the container
func (c *Container) Shutdown() error {
	return nil
}

// initConfig initializes the configuration
func (c *Container) initConfig() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web server
func (c *Container) initWeb() {
	c.Web = echo.New()

	switch c.Config.App.Environment {
	case config.EnvProduction:
		c.Web.Logger.SetLevel(log.WARN)
	default:
		c.Web.Logger.SetLevel(log.DEBUG)
	}

	c.Web.Validator = c.Validator
}

// initTemplateRenderer initializes the template renderer
func (c *Container) initTemplateRenderer() {
	c.TemplateRenderer = NewTemplateRenderer(c.Config)
}
