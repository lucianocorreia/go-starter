package handlers

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/lucianocorreia/go-starter/database/sqlc"
)

// Handlers is the application handlers struct
type Handlers struct {
	store database.Store
}

// NewHandlers returns a new Handlers struct
func NewHandlers(store database.Store) *Handlers {
	return &Handlers{
		store: store,
	}
}

func (h *Handlers) RenderPage(c *fiber.Ctx, page string, title string, layouts ...string) error {
	return c.Render(page, fiber.Map{"Title": title}, layouts...)
}
