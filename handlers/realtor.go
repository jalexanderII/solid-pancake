package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/middleware"

	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/models"
)

type Realtor struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func CreateResponseRealtor(realtorModel models.Realtor) Realtor {
	return Realtor{ID: realtorModel.ID, Name: realtorModel.Name, Company: realtorModel.Company, PhoneNumber: realtorModel.PhoneNumber}
}

func findRealtor(id int, realtor *models.Realtor) error {
	database.Database.Db.Find(&realtor, "id = ?", id)
	if realtor.ID == 0 {
		return errors.New("realtor does not exist")
	}
	return nil
}

func CreateRealtor(c *fiber.Ctx) error {
	var realtor models.Realtor
	if err := c.BodyParser(&realtor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}
	responseRealtor := CreateResponseRealtor(realtor)
	errs := middleware.ValidateStruct(&responseRealtor)
	if errs != nil {
		return c.JSON(errs)
	}
	if err := database.Database.Db.Create(&realtor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create realtor", "data": err.Error()})
	}
	responseRealtor.ID = realtor.ID

	return c.Status(fiber.StatusOK).JSON(responseRealtor)
}

func GetRealtors(c *fiber.Ctx) error {
	var realtors []models.Realtor
	database.Database.Db.Find(&realtors)

	responseRealtors := make([]Realtor, len(realtors))
	for idx, realtor := range realtors {
		responseRealtors[idx] = CreateResponseRealtor(realtor)
	}

	return c.Status(fiber.StatusOK).JSON(responseRealtors)
}
