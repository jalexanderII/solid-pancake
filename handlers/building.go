package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	"github.com/jalexanderII/solid-pancake/models"
)

type Building struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Address     models.Place `json:"address" validate:"dive"`
	Description string       `json:"description,omitempty"`
	Amenities   []string     `json:"amenities"`
	Realtor     Realtor      `json:"realtor" validate:"dive"`
}

func CreateResponseBuilding(buildingModel models.Building, realtor Realtor) Building {
	return Building{
		ID:          buildingModel.ID,
		Name:        buildingModel.Name,
		Address:     buildingModel.Address,
		Description: buildingModel.Description,
		Amenities:   buildingModel.Amenities,
		Realtor:     realtor,
	}
}

func findBuilding(id int, building *models.Building) error {
	database.Database.Db.Find(&building, "id = ?", id)
	if building.ID == 0 {
		return errors.New("building does not exist")
	}
	return nil
}

func CreateBuilding(c *fiber.Ctx) error {
	var building models.Building
	var realtor models.Realtor

	if err := c.BodyParser(&building); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}
	if err := findRealtor(building.RealtorRef, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	realtorResponse := CreateResponseRealtor(realtor)
	responseBuilding := CreateResponseBuilding(building, realtorResponse)
	errs := middleware.ValidateStruct(&responseBuilding)
	if errs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": errs})
	}

	if err := database.Database.Db.Create(&building).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create building", "data": err.Error()})
	}
	responseBuilding.ID = building.ID
	return c.Status(fiber.StatusOK).JSON(responseBuilding)
}

func GetBuildings(c *fiber.Ctx) error {
	var buildings []models.Building
	database.Database.Db.Find(&buildings)

	responseBuildings := make([]Building, len(buildings))
	for idx, building := range buildings {
		var realtor models.Realtor
		database.Database.Db.Find(&realtor, "id = ?", building.RealtorRef)
		realtorResponse := CreateResponseRealtor(realtor)
		responseBuildings[idx] = CreateResponseBuilding(building, realtorResponse)
	}
	return c.Status(fiber.StatusOK).JSON(responseBuildings)
}
