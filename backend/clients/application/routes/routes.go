package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/clients/application/handlers"
)

func SetupApplicationRoutes(v1 fiber.Router) {
	// Application endpoints
	application := v1.Group("/application")
	application.Post("/apply", handlers.Apply)
	application.Post("/:id/upload", handlers.Upload)
	// Application Request endpoints
	appReq := application.Group("/request")
	appReq.Get("/", handlers.GetApplications)
	appReq.Get("/:id", handlers.GetApplication)
	appReq.Delete("/:id", handlers.DeleteApplication)
	// Application Response endpoints
	appResp := application.Group("/response")
	appResp.Get("/:id", handlers.GetApplicationResponse)
	appResp.Delete("/:id", handlers.DeleteApplicationResponse)
}