package routes

import (
	"github.com/gofiber/fiber/v2"
	RealEstateH "github.com/jalexanderII/solid-pancake/services/realestate/handlers"
)

func SetupRealEstateRoutes(v1 fiber.Router) {
	// Realtor endpoints
	realtors := v1.Group("/realtors")
	realtors.Get("/", RealEstateH.GetRealtors)
	realtors.Post("/", RealEstateH.CreateRealtor)
	realtors.Get("/:id", RealEstateH.GetRealtor)
	realtors.Patch("/:id", RealEstateH.UpdateRealtor)
	realtors.Delete("/:id", RealEstateH.DeleteRealtor)

	// Building endpoints
	buildings := v1.Group("/buildings")
	buildings.Get("/", RealEstateH.GetBuildings)
	buildings.Post("/", RealEstateH.CreateBuilding)
	buildings.Get("/:id", RealEstateH.GetBuilding)
	buildings.Patch("/:id", RealEstateH.UpdateBuilding)
	buildings.Patch("/:id/realtor/:realtor_id", RealEstateH.UpdateBuildingRealtor)
	buildings.Delete("/:id", RealEstateH.DeleteBuilding)

	// Apartment endpoints
	apartments := v1.Group("/apartments")
	apartments.Get("/", RealEstateH.GetApartments)
	apartments.Post("/", RealEstateH.CreateApartment)
	apartments.Get("/:id", RealEstateH.GetApartment)
	apartments.Patch("/:id", RealEstateH.UpdateApartment)
	apartments.Patch("/:id/building/:building_id", RealEstateH.UpdateApartmentBuilding)
	apartments.Patch("/:id/realtor/:realtor_id", RealEstateH.UpdateApartmentRealtor)
	apartments.Delete("/:id", RealEstateH.DeleteApartment)
}
