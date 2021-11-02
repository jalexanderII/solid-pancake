package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	"github.com/jalexanderII/solid-pancake/models"
)

// Apartment To be used as a serializer
type Apartment struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	Address        models.Place          `json:"address" validate:"dive"`
	Rent           int                   `json:"rent"`
	Size           float32               `json:"size"`
	Features       models.Features       `json:"features" validate:"dive"`
	ListingMetrics models.ListingMetrics `json:"listing_metrics" validate:"dive"`
	Description    string                `json:"description"`
	Amenities      []string              `json:"amenities"`
	Building       Building              `json:"building" validate:"dive"`
	Realtor        Realtor               `json:"realtor" validate:"dive"`
}

// CreateResponseApartment Takes in a model and returns a serializer
func CreateResponseApartment(apartmentModel models.Apartment, building Building) Apartment {
	return Apartment{
		ID:             apartmentModel.ID,
		Name:           apartmentModel.Name,
		Address:        apartmentModel.Address,
		Rent:           apartmentModel.Rent,
		Size:           apartmentModel.Size,
		Features:       apartmentModel.Features,
		ListingMetrics: apartmentModel.ListingMetrics,
		Description:    apartmentModel.Description,
		Amenities:      apartmentModel.Amenities,
		Building:       building,
		Realtor:        building.Realtor,
	}
}

func CreateApartment(c *fiber.Ctx) error {
	var apartment models.Apartment
	var building models.Building
	var realtor models.Realtor

	if err := c.BodyParser(&apartment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}
	if err := findBuilding(apartment.BuildingRef, &building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := findRealtor(apartment.RealtorRef, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	realtorResponse := CreateResponseRealtor(realtor)
	responseApartment := CreateResponseApartment(
		apartment,
		CreateResponseBuilding(building, realtorResponse),
	)

	errs := middleware.ValidateStruct(&responseApartment)
	if errs != nil {
		return c.JSON(errs)
	}

	if err := database.Database.Db.Create(&apartment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create apartment", "data": err.Error()})
	}
	responseApartment.ID = apartment.ID
	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

func GetApartments(c *fiber.Ctx) error {
	var apartments []models.Apartment
	database.Database.Db.Find(&apartments)
	responseApartments := make([]Apartment, len(apartments))

	for idx, apartment := range apartments {
		var realtor models.Realtor
		if err := findRealtor(apartment.RealtorRef, &realtor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		realtorResponse := CreateResponseRealtor(realtor)

		var building models.Building
		if err := findBuilding(apartment.BuildingRef, &building); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		buildingResponse := CreateResponseBuilding(building, realtorResponse)

		responseApartments[idx] = CreateResponseApartment(apartment, buildingResponse)
	}
	return c.Status(fiber.StatusOK).JSON(responseApartments)
}

func findApartment(id int, apartment *models.Apartment) error {
	database.Database.Db.Find(&apartment, "id = ?", id)
	if apartment.ID == 0 {
		return errors.New("apartment does not exist")
	}
	return nil
}

func GetApartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var apartment models.Apartment
	if err := findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var building models.Building
	var realtor models.Realtor
	database.Database.Db.First(&realtor, apartment.RealtorRef)
	realtorResponse := CreateResponseRealtor(realtor)

	database.Database.Db.First(&building, apartment.BuildingRef)
	responseApartment := CreateResponseApartment(
		apartment,
		CreateResponseBuilding(building, realtorResponse),
	)

	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

type UpdateApartmentResponse struct {
	Name           string                `json:"name"`
	Address        models.Place          `json:"address" validate:"dive"`
	Rent           int                   `json:"rent"`
	Size           float32               `json:"size"`
	Features       models.Features       `json:"features" validate:"dive"`
	ListingMetrics models.ListingMetrics `json:"listing_metrics" validate:"dive"`
	Description    string                `json:"description"`
	Amenities      []string              `json:"amenities"`
}

func UpdateApartment(c *fiber.Ctx) error {
	var apartment models.Apartment
	var updateApartmentResponse UpdateApartmentResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	if err = findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err = c.BodyParser(&updateApartmentResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	apartment.Name = updateApartmentResponse.Name
	apartment.Address = updateApartmentResponse.Address
	apartment.Rent = updateApartmentResponse.Rent
	apartment.Size = updateApartmentResponse.Size
	apartment.Features = updateApartmentResponse.Features
	apartment.ListingMetrics = updateApartmentResponse.ListingMetrics
	apartment.Description = updateApartmentResponse.Description
	apartment.Amenities = updateApartmentResponse.Amenities
	database.Database.Db.Save(&apartment)

	var building models.Building
	var realtor models.Realtor
	database.Database.Db.First(&building, apartment.BuildingRef)
	database.Database.Db.First(&realtor, apartment.RealtorRef)
	realtorResponse := CreateResponseRealtor(realtor)
	responseApartment := CreateResponseApartment(
		apartment,
		CreateResponseBuilding(building, realtorResponse),
	)
	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

func DeleteApartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var apartment models.Apartment

	database.Database.Db.First(&apartment, id)
	if apartment.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No apartment found with ID", "data": nil})
	}
	database.Database.Db.Delete(&apartment)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Apartment successfully deleted", "data": nil})
}
