package main

import "github.com/gofiber/fiber/v2"

func (a *App) setupRoutes(app *fiber.App) {
	app.Get("/foo", a.Handlers.Home)
}
