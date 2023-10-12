package handlers

import (
	"os"
	"testing"

	"github.com/lucianocorreia/go-starter/config"
	"github.com/lucianocorreia/go-starter/pkg/services"
)

var (
	c *services.Container
)

func TestMain(m *testing.M) {
	config.SwitchEnvironment(config.EnvTest)

	c = services.NewContainer()

	exitVal := m.Run()

	if err := c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
