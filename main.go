package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/config"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	ApplicationRoutes "github.com/jalexanderII/solid-pancake/services/application/routes"
	RestRoutes "github.com/jalexanderII/solid-pancake/services/realestate/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	middleware.FiberMiddleware(app)

	v1 := config.SetupV1Routes(app)
	RestRoutes.SetupRealEstateRoutes(v1)
	ApplicationRoutes.SetupApplicationRoutes(v1)

	// Start server (with graceful shutdown).
	config.StartServerWithGracefulShutdown(app)
}
