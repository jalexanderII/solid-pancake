package routes

import (
	"github.com/gofiber/fiber/v2"
	ApplicationH "github.com/jalexanderII/solid-pancake/services/application/handlers"
)

func SetupApplicationRoutes(v1 fiber.Router) {
	// Application endpoints
	application := v1.Group("/application")
	application.Post("/apply", ApplicationH.Apply)
	application.Delete("/:id", ApplicationH.DeleteApplication)
}
