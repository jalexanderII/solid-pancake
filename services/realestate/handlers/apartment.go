package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	"github.com/jalexanderII/solid-pancake/utils"
)

// Apartment To be used as a serializer
type Apartment struct {
	ID             uint                       `json:"id"`
	Name           string                     `json:"name"`
	Address        RealEstateM.Place          `json:"address" validate:"dive"`
	Rent           int                        `json:"rent"`
	Size           float32                    `json:"size"`
	Features       RealEstateM.Features       `json:"features" validate:"dive"`
	ListingMetrics RealEstateM.ListingMetrics `json:"listing_metrics" validate:"dive"`
	Description    string                     `json:"description"`
	Amenities      []string                   `json:"amenities"`
	Images         []string                   `json:"images"`
	Building       Building                   `json:"building" validate:"dive"`
	Realtor        Realtor                    `json:"realtor" validate:"dive"`
}

// CreateResponseApartment Takes in a model and returns a serializer
func CreateResponseApartment(apartmentModel RealEstateM.Apartment) Apartment {
	var building RealEstateM.Building
	var realtor RealEstateM.Realtor
	database.Database.Db.First(&building, apartmentModel.BuildingRef)
	database.Database.Db.First(&realtor, building.RealtorRef)
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
		Images:         apartmentModel.Images,
		Building:       CreateResponseBuilding(building),
		Realtor:        CreateResponseRealtor(realtor),
	}
}

func CreateApartment(c *fiber.Ctx) error {
	var apartment RealEstateM.Apartment

	if err := c.BodyParser(&apartment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}
	responseApartment := CreateResponseApartment(apartment)

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
	var apartments []RealEstateM.Apartment
	database.Database.Db.Find(&apartments)
	responseApartments := make([]Apartment, len(apartments))

	for idx, apartment := range apartments {
		responseApartments[idx] = CreateResponseApartment(apartment)
	}
	return c.Status(fiber.StatusOK).JSON(responseApartments)
}

func findApartment(id int, apartment *RealEstateM.Apartment) error {
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

	var apartment RealEstateM.Apartment
	if err := findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	responseApartment := CreateResponseApartment(apartment)

	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

type UpdateApartmentResponse struct {
	Name           string                     `json:"name"`
	Address        RealEstateM.Place          `json:"address" validate:"dive"`
	Rent           int                        `json:"rent"`
	Size           float32                    `json:"size"`
	Features       RealEstateM.Features       `json:"features" validate:"dive"`
	ListingMetrics RealEstateM.ListingMetrics `json:"listing_metrics" validate:"dive"`
	Description    string                     `json:"description"`
	Amenities      []string                   `json:"amenities"`
	Images         []string                   `json:"images"`
}

func UpdateApartment(c *fiber.Ctx) error {
	var apartment RealEstateM.Apartment
	var uar UpdateApartmentResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	if err = findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err = c.BodyParser(&uar); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	apartment.Name = utils.UpdateIfNew(uar.Name, apartment.Name).(string)
	apartment.Address = utils.UpdateIfNew(uar.Address, apartment.Address).(RealEstateM.Place)
	apartment.Rent = utils.UpdateIfNew(uar.Rent, apartment.Rent).(int)
	apartment.Size = utils.UpdateIfNew(uar.Size, apartment.Size).(float32)
	apartment.Features = utils.UpdateIfNew(uar.Features, apartment.Features).(RealEstateM.Features)
	apartment.ListingMetrics = utils.UpdateIfNew(uar.ListingMetrics, apartment.ListingMetrics).(RealEstateM.ListingMetrics)
	apartment.Description = utils.UpdateIfNew(uar.Description, apartment.Description).(string)
	apartment.Amenities = utils.UpdateIfNew(uar.Amenities, apartment.Amenities).([]string)
	apartment.Images = utils.UpdateIfNew(uar.Images, apartment.Images).([]string)
	database.Database.Db.Save(&apartment)

	responseApartment := CreateResponseApartment(apartment)
	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

func UpdateApartmentBuilding(c *fiber.Ctx) error {
	buildingID, err := c.ParamsInt("building_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var building RealEstateM.Building
	if err := findBuilding(buildingID, &building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var apartment RealEstateM.Apartment
	if err = findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	apartment.Building = building
	database.Database.Db.Save(&apartment)

	responseApartment := CreateResponseApartment(apartment)
	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

func UpdateApartmentRealtor(c *fiber.Ctx) error {
	realtorID, err := c.ParamsInt("realtor_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var realtor RealEstateM.Realtor
	if err := findRealtor(realtorID, &realtor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var apartment RealEstateM.Apartment
	if err = findApartment(id, &apartment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	apartment.Realtor = realtor
	database.Database.Db.Save(&apartment)

	UpdateBuildingRealtorUrl := fmt.Sprint("http://127.0.0.1:9092/api/v1/buildings/", id, "/realtor/", realtorID)
	_, err = utils.MakeSyncPatchCall(UpdateBuildingRealtorUrl, realtor, "PATCH")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	responseApartment := CreateResponseApartment(apartment)
	return c.Status(fiber.StatusOK).JSON(responseApartment)
}

func DeleteApartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var apartment RealEstateM.Apartment

	database.Database.Db.First(&apartment, id)
	if apartment.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No apartment found with ID"})
	}
	database.Database.Db.Delete(&apartment)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Apartment successfully deleted"})
}
