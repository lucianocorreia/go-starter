package main

import "github.com/gofiber/fiber/v2"

func (a *App) setupRoutes(app *fiber.App) {
	app.Get("/", a.Handlers.Home)

	// auth routes
	app.Get("/register", a.Handlers.Register)
}
