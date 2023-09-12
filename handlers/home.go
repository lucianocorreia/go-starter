package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) Home(c *fiber.Ctx) error {
	return h.RenderPage(c, "home", "Home")
}
