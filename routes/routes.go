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
	realtors.Get("/:id", handlers.GetRealtor)
	realtors.Patch("/:id", handlers.UpdateRealtor)
	realtors.Delete("/:id", handlers.DeleteRealtor)

	// Building endpoints
	buildings := v1.Group("/buildings")
	buildings.Get("/", handlers.GetBuildings)
	buildings.Post("/", handlers.CreateBuilding)
	buildings.Get("/:id", handlers.GetBuilding)
	buildings.Patch("/:id", handlers.UpdateBuilding)
	buildings.Patch("/:id/realtor/:realtor_id", handlers.UpdateBuildingRealtor)
	buildings.Delete("/:id", handlers.DeleteBuilding)

	// Apartment endpoints
	apartments := v1.Group("/apartments")
	apartments.Get("/", handlers.GetApartments)
	apartments.Post("/", handlers.CreateApartment)
	apartments.Get("/:id", handlers.GetApartment)
	apartments.Patch("/:id", handlers.UpdateApartment)
	apartments.Patch("/:id/building/:building_id", handlers.UpdateApartmentBuilding)
	apartments.Patch("/:id/realtor/:realtor_id", handlers.UpdateApartmentRealtor)
	apartments.Delete("/:id", handlers.DeleteApartment)
}
