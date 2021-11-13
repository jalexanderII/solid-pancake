package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	LifeCycleM "github.com/jalexanderII/solid-pancake/services/lifecycle/models"
	UserH "github.com/jalexanderII/solid-pancake/services/users/handlers"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
)

// PaymentRequest To be used as a serializer
type PaymentRequest struct {
	ID      uint       `json:"id"`
	Purpose string     `json:"purpose" validate:"required"`
	Amount  float32    `json:"amount" validate:"required"`
	Period  string     `json:"period"`
	User    UserH.User `json:"user" validate:"dive"`
}

// CreatePaymentRequest Takes in a model and returns a serializer
func CreatePaymentRequest(paymentRequest LifeCycleM.PaymentRequest) PaymentRequest {
	var user UserM.User
	database.Database.Db.First(&user, paymentRequest.UserRef)
	return PaymentRequest{
		ID:      paymentRequest.ID,
		Purpose: paymentRequest.Purpose,
		Amount:  paymentRequest.Amount,
		Period:  paymentRequest.Period,
		User:    UserH.CreateResponseUser(user),
	}
}

func RequestPayment(c *fiber.Ctx) error {
	var paymentRequest LifeCycleM.PaymentRequest
	if err := c.BodyParser(&paymentRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}

	if err := database.Database.Db.Create(&paymentRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create payment request", "data": err.Error()})
	}
	paymentRequestResponse := CreatePaymentRequest(paymentRequest)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "submitted payment request", "user": paymentRequestResponse.User.Name})
}
