package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) Register(c *fiber.Ctx) error {
	return h.RenderPage(c, "auth/register", "Register", "layouts/auth")
}
