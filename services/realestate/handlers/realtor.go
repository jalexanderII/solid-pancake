package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	"gorm.io/gorm/clause"
)

type Realtor struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func CreateResponseRealtor(realtorModel RealEstateM.Realtor) Realtor {
	return Realtor{ID: realtorModel.ID, Name: realtorModel.Name, Company: realtorModel.Company, PhoneNumber: realtorModel.PhoneNumber}
}

func CreateRealtor(c *fiber.Ctx) error {
	var realtor RealEstateM.Realtor
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
	var realtors []RealEstateM.Realtor
	database.Database.Db.Find(&realtors)

	responseRealtors := make([]Realtor, len(realtors))
	for idx, realtor := range realtors {
		responseRealtors[idx] = CreateResponseRealtor(realtor)
	}

	return c.Status(fiber.StatusOK).JSON(responseRealtors)
}

func findRealtor(id int, realtor *RealEstateM.Realtor) error {
	database.Database.Db.Find(&realtor, "id = ?", id)
	if realtor.ID == 0 {
		return errors.New("realtor does not exist")
	}
	return nil
}

func GetRealtor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var realtor RealEstateM.Realtor
	if err := findRealtor(id, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseRealtor := CreateResponseRealtor(realtor)
	return c.Status(fiber.StatusOK).JSON(responseRealtor)
}

type UpdateRealtorResponse struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number"`
}

func UpdateRealtor(c *fiber.Ctx) error {
	var realtor RealEstateM.Realtor
	var urr UpdateRealtorResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	if err = findRealtor(id, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err = c.BodyParser(&urr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	database.Database.Db.Model(&realtor).Clauses(clause.Returning{}).Updates(urr)

	responseRealtor := CreateResponseRealtor(realtor)
	return c.Status(fiber.StatusOK).JSON(responseRealtor)
}

func DeleteRealtor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var realtor RealEstateM.Realtor

	database.Database.Db.First(&realtor, id)
	if realtor.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No realtor found with ID"})
	}
	database.Database.Db.Delete(&realtor)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Realtor successfully deleted"})
}
