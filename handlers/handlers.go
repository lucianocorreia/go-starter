package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	database "github.com/lucianocorreia/go-starter/database/sqlc"
)

// Handlers is the application handlers struct
type Handlers struct {
	store   database.Store
	session *session.Store
}

// NewHandlers returns a new Handlers struct
func NewHandlers(store database.Store, session *session.Store) *Handlers {
	return &Handlers{
		store:   store,
		session: session,
	}
}

// RenderPage renders a page
func (h *Handlers) RenderPage(c *fiber.Ctx, page string, title string, layouts ...string) error {
	return c.Render(page, fiber.Map{"Title": title}, layouts...)
}
