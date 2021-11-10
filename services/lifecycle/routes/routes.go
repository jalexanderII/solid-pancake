package routes

import (
	"github.com/gofiber/fiber/v2"
	LifeCycleH "github.com/jalexanderII/solid-pancake/services/lifecycle/handlers"
)

func SetupLifeCycleRoutes(v1 fiber.Router) {
	// Payments endpoints
	payments := v1.Group("/payments")

	// Payment request endpoints
	request := payments.Group("/request")
	request.Post("/", LifeCycleH.RequestPayment)

	// Payment response endpoints
	pay := payments.Group("/pay")
	pay.Post("/", LifeCycleH.Pay)
}
