package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/config"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	"github.com/jalexanderII/solid-pancake/services/realestate/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	middleware.FiberMiddleware(app)

	v1 := config.SetupV1Routes(app)
	routes.SetupRealEstateRoutes(v1)

	// Start server (with graceful shutdown).
	config.StartServerWithGracefulShutdown(app)
}
