package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	"github.com/jalexanderII/solid-pancake/models"
	"github.com/jalexanderII/solid-pancake/utils"
)

type Building struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Address     models.Place `json:"address"`
	Description string       `json:"description"`
	Amenities   []string     `json:"amenities"`
	Images      []string     `json:"images"`
	Realtor     Realtor      `json:"realtor" validate:"dive"`
}

func CreateResponseBuilding(buildingModel models.Building) Building {
	var realtor models.Realtor
	database.Database.Db.First(&realtor, buildingModel.RealtorRef)
	return Building{
		ID:          buildingModel.ID,
		Name:        buildingModel.Name,
		Address:     buildingModel.Address,
		Description: buildingModel.Description,
		Amenities:   buildingModel.Amenities,
		Images:      buildingModel.Images,
		Realtor:     CreateResponseRealtor(realtor),
	}
}

func CreateBuilding(c *fiber.Ctx) error {
	var building models.Building

	if err := c.BodyParser(&building); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}
	responseBuilding := CreateResponseBuilding(building)
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
		responseBuildings[idx] = CreateResponseBuilding(building)
	}
	return c.Status(fiber.StatusOK).JSON(responseBuildings)
}

func findBuilding(id int, building *models.Building) error {
	database.Database.Db.Find(&building, "id = ?", id)
	if building.ID == 0 {
		return errors.New("building does not exist")
	}
	return nil
}

func GetBuilding(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var building models.Building
	if err := findBuilding(id, &building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	responseBuilding := CreateResponseBuilding(building)

	return c.Status(fiber.StatusOK).JSON(responseBuilding)
}

type UpdateBuildingResponse struct {
	Name        string       `json:"name"`
	Address     models.Place `json:"address"`
	Description string       `json:"description"`
	Amenities   []string     `json:"amenities"`
	Images      []string     `json:"images"`
}

func UpdateBuilding(c *fiber.Ctx) error {
	var building models.Building
	var ubr UpdateBuildingResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	if err = findBuilding(id, &building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err = c.BodyParser(&ubr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	building.Name = utils.UpdateIfNew(ubr.Name, building.Name).(string)
	building.Address = utils.UpdateIfNew(ubr.Address, building.Address).(models.Place)
	building.Description = utils.UpdateIfNew(ubr.Description, building.Description).(string)
	building.Amenities = utils.UpdateIfNew(ubr.Amenities, building.Amenities).([]string)
	building.Images = utils.UpdateIfNew(ubr.Images, building.Images).([]string)
	database.Database.Db.Save(&building)

	responseBuilding := CreateResponseBuilding(building)

	return c.Status(fiber.StatusOK).JSON(responseBuilding)
}

func UpdateBuildingRealtor(c *fiber.Ctx) error {
	realtorID, err := c.ParamsInt("realtor_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var realtor models.Realtor
	if err := findRealtor(realtorID, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var building models.Building
	if err = findBuilding(id, &building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	building.Realtor = realtor
	database.Database.Db.Save(&building)

	responseBuilding := CreateResponseBuilding(building)
	return c.Status(fiber.StatusOK).JSON(responseBuilding)
}

func DeleteBuilding(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var building models.Building

	database.Database.Db.First(&building, id)
	if building.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No building found with ID"})
	}
	database.Database.Db.Delete(&building)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Building successfully deleted"})
}
