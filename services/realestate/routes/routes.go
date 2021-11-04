package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers2 "github.com/jalexanderII/solid-pancake/services/realestate/handlers"
)

func SetupRealEstateRoutes(v1 fiber.Router) {
	// Realtor endpoints
	realtors := v1.Group("/realtors")
	realtors.Get("/", handlers2.GetRealtors)
	realtors.Post("/", handlers2.CreateRealtor)
	realtors.Get("/:id", handlers2.GetRealtor)
	realtors.Patch("/:id", handlers2.UpdateRealtor)
	realtors.Delete("/:id", handlers2.DeleteRealtor)

	// Building endpoints
	buildings := v1.Group("/buildings")
	buildings.Get("/", handlers2.GetBuildings)
	buildings.Post("/", handlers2.CreateBuilding)
	buildings.Get("/:id", handlers2.GetBuilding)
	buildings.Patch("/:id", handlers2.UpdateBuilding)
	buildings.Patch("/:id/realtor/:realtor_id", handlers2.UpdateBuildingRealtor)
	buildings.Delete("/:id", handlers2.DeleteBuilding)

	// Apartment endpoints
	apartments := v1.Group("/apartments")
	apartments.Get("/", handlers2.GetApartments)
	apartments.Post("/", handlers2.CreateApartment)
	apartments.Get("/:id", handlers2.GetApartment)
	apartments.Patch("/:id", handlers2.UpdateApartment)
	apartments.Patch("/:id/building/:building_id", handlers2.UpdateApartmentBuilding)
	apartments.Patch("/:id/realtor/:realtor_id", handlers2.UpdateApartmentRealtor)
	apartments.Delete("/:id", handlers2.DeleteApartment)
}
