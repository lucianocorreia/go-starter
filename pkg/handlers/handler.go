package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/pkg/services"
)

// Handler is the struct that holds all handlers
type Handler struct {
	Container *services.Container
}

// NewHandler creates a new handler
func NewHandler(c *services.Container) *Handler {
	return &Handler{
		Container: c,
	}
}

// RenderPage renders a page
func (h *Handler) RenderPage(ctx echo.Context, page Page) error {
	return nil
}
