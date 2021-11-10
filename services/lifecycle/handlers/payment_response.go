package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	LifeCycleM "github.com/jalexanderII/solid-pancake/services/lifecycle/models"
	RealEstateH "github.com/jalexanderII/solid-pancake/services/realestate/handlers"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	UserH "github.com/jalexanderII/solid-pancake/services/users/handlers"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
)

// PaymentResponse To be used as a serializer
type PaymentResponse struct {
	ID             uint                   `json:"id"`
	Payment        LifeCycleM.PaymentInfo `json:"payment" validate:"dive"`
	User           UserH.User             `json:"user" validate:"dive"`
	Apartment      RealEstateH.Apartment  `json:"apartment" validate:"dive"`
	PaymentRequest PaymentRequest         `json:"payment_request" validate:"dive"`
}

// CreatePaymentResponse Takes in a model and returns a serializer
func CreatePaymentResponse(paymentResponse LifeCycleM.PaymentResponse) PaymentResponse {
	var (
		user           UserM.User
		apartment      RealEstateM.Apartment
		paymentRequest LifeCycleM.PaymentRequest
	)
	database.Database.Db.First(&user, paymentResponse.UserRef)
	database.Database.Db.First(&apartment, paymentResponse.ApartmentRef)
	database.Database.Db.First(&paymentRequest, paymentResponse.PaymentRef)

	return PaymentResponse{
		ID:             paymentResponse.ID,
		Payment:        paymentResponse.Payment,
		User:           UserH.CreateResponseUser(user),
		Apartment:      RealEstateH.CreateResponseApartment(apartment),
		PaymentRequest: CreatePaymentRequest(paymentRequest),
	}
}

func findPaymentResponse(id int, paymentResponse *LifeCycleM.PaymentResponse) error {
	database.Database.Db.Find(&paymentResponse, "id = ?", id)
	if paymentResponse.ID == 0 {
		return errors.New("payment response does not exist")
	}
	return nil
}

func Pay(c *fiber.Ctx) error {
	var paymentResponse LifeCycleM.PaymentResponse
	if err := c.BodyParser(&paymentResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}

	if err := database.Database.Db.Create(&paymentResponse).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create payment response", "data": err.Error()})
	}
	paymentConfirmationResponse, err := SendToPaymentService(int(paymentResponse.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error with payment process", "data": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(paymentConfirmationResponse)
}
