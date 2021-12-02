package main

import (
	"github.com/gofiber/fiber/v2"
	ApplicationRoutes "github.com/jalexanderII/solid-pancake/clients/application/routes"
	"github.com/jalexanderII/solid-pancake/config"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	LifeCycleRoutes "github.com/jalexanderII/solid-pancake/services/lifecycle/routes"
	RestRoutes "github.com/jalexanderII/solid-pancake/services/realestate/routes"
	UserRoutes "github.com/jalexanderII/solid-pancake/services/users/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	middleware.FiberMiddleware(app)

	v1 := config.SetupV1Routes(app)
	RestRoutes.SetupRealEstateRoutes(v1)
	ApplicationRoutes.SetupApplicationRoutes(v1)
	UserRoutes.SetupUserAndAuthRoutes(v1)
	LifeCycleRoutes.SetupLifeCycleRoutes(v1)

	// Start server (with graceful shutdown).
	config.StartServerWithGracefulShutdown(app)
}
