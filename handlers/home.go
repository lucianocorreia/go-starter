package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}
