package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) Register(c *fiber.Ctx) error {

	return h.RenderPage(c, "auth/register", "Register", "layouts/auth")
}

func (h *Handlers) RegisterPost(c *fiber.Ctx) error {
	fmt.Println("RegisterPost")

	return c.Redirect("/")
}
