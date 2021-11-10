package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/middleware"
	UserH "github.com/jalexanderII/solid-pancake/services/users/handlers"
)

func SetupUserAndAuthRoutes(v1 fiber.Router) {
	// Auth
	auth := v1.Group("/auth")
	auth.Post("/register", UserH.Register)
	auth.Post("/login", UserH.Login)
	auth.Post("/logout", UserH.Logout)

	// User endpoints
	users := v1.Group("/users")
	users.Get("/", UserH.GetUsers)
	users.Get("/:id", UserH.GetUser)
	users.Patch("/:id", middleware.Protected(), UserH.UpdateUser)
	users.Delete("/:id", middleware.Protected(), UserH.DeleteUser)
}
