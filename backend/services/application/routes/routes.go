package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/services/application/client"
)

func SetupApplicationRoutes(v1 fiber.Router) {
	// Application endpoints
	application := v1.Group("/application")
	application.Post("/apply", client.Apply)
	application.Post("/:id/upload", client.Upload)
	// Application Request endpoints
	appReq := application.Group("/request")
	appReq.Get("/", client.GetApplications)
	appReq.Get("/:id", client.GetApplication)
	appReq.Delete("/:id", client.DeleteApplication)
	// Application Response endpoints
	appResp := application.Group("/response")
	appResp.Get("/:id", client.GetApplicationResponse)
	appResp.Delete("/:id", client.DeleteApplicationResponse)
}
