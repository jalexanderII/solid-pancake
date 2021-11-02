package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/jalexanderII/solid-pancake/handlers"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// monitoring api stats
	v1.Get("/dashboard", monitor.New())

	// Realtor endpoints
	realtors := v1.Group("/realtors")
	realtors.Get("/", handlers.GetRealtors)
	realtors.Post("/", handlers.CreateRealtor)

	// Building endpoints
	buildings := v1.Group("/buildings")
	buildings.Get("/", handlers.GetBuildings)
	buildings.Post("/", handlers.CreateBuilding)

	// Apartment endpoints
	apartments := v1.Group("/apartments")
	apartments.Get("/", handlers.GetApartments)
	apartments.Get("/:id", handlers.GetApartment)
	apartments.Post("/", handlers.CreateApartment)
	apartments.Patch("/:id", handlers.UpdateApartment)
	apartments.Delete("/:id", handlers.DeleteApartment)
}
