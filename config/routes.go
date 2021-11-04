package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupV1Routes(app *fiber.App) fiber.Router {
	app.Get("/", welcome)
	api := app.Group("/api")
	v1 := api.Group("/v1")
	// monitoring api stats
	v1.Get("/dashboard", monitor.New())
	return v1
}
