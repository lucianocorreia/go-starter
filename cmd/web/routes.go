package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/utils"
)

func (a *App) setupRoutes(app *fiber.App) {
	app.Get("/", a.Handlers.Home)

	// auth routes
	app.Get("/register", csrfProtection, a.Handlers.Register)
	app.Post("/register", a.Handlers.RegisterPost)
}

// Add CSRF protection middleware.
// Should be done AFTER session middleware.
var csrfProtection = csrf.New(csrf.Config{
	KeyLookup:      "form:_csrf",
	CookieName:     "csrf_",
	CookieSameSite: "Strict",
	Expiration:     1 * time.Hour,
	KeyGenerator:   utils.UUID,
	ContextKey:     "token",
})
