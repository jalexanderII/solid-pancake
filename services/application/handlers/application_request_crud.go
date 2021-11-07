package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
	RealEstateH "github.com/jalexanderII/solid-pancake/services/realestate/handlers"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
)

type ApplicantFormRequest struct {
	ID              uint                  `json:"id"`
	Name            string                `json:"name"`
	SocialSecurity  string                `json:"social_security,omitempty"`
	DateOfBirth     string                `json:"date_of_birth,omitempty"`
	DriversLicense  string                `json:"drivers_license,omitempty"`
	PreviousAddress RealEstateM.Place     `json:"previous_address,omitempty" validate:"dive"`
	Landlord        string                `json:"landlord,omitempty"`
	LandlordNumber  string                `json:"landlord_number,omitempty"`
	Employer        string                `json:"employer,omitempty"`
	Salary          int32                 `json:"salary,omitempty"`
	Apartment       RealEstateH.Apartment `json:"apartment" validate:"dive"`
}

// CreateApplicantFormRequest Takes in a model and returns a serializer
func CreateApplicantFormRequest(applicantRequestModel ApplicationM.ApplicantFormRequest) ApplicantFormRequest {
	var apartment RealEstateM.Apartment
	database.Database.Db.First(&apartment, applicantRequestModel.ApartmentRef)
	return ApplicantFormRequest{
		ID:              applicantRequestModel.ID,
		Name:            applicantRequestModel.Name,
		SocialSecurity:  applicantRequestModel.SocialSecurity,
		DateOfBirth:     applicantRequestModel.DateOfBirth,
		DriversLicense:  applicantRequestModel.DriversLicense,
		PreviousAddress: applicantRequestModel.PreviousAddress,
		Landlord:        applicantRequestModel.Landlord,
		LandlordNumber:  applicantRequestModel.LandlordNumber,
		Employer:        applicantRequestModel.Employer,
		Salary:          applicantRequestModel.Salary,
		Apartment:       RealEstateH.CreateResponseApartment(apartment),
	}
}

func GetApplications(c *fiber.Ctx) error {
	var applications []ApplicationM.ApplicantFormRequest
	database.Database.Db.Find(&applications)

	responseApplRequests := make([]ApplicantFormRequest, len(applications))
	for idx, application := range applications {
		responseApplRequests[idx] = CreateApplicantFormRequest(application)
	}
	return c.Status(fiber.StatusOK).JSON(responseApplRequests)
}

func findApplication(id int, application *ApplicationM.ApplicantFormRequest) error {
	database.Database.Db.Find(&application, "id = ?", id)
	if application.ID == 0 {
		return errors.New("application does not exist")
	}
	return nil
}

func GetApplication(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var application ApplicationM.ApplicantFormRequest
	database.Database.Db.First(&application, id)
	if application.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found with ID"})
	}

	responseApplRequest := CreateApplicantFormRequest(application)

	return c.Status(fiber.StatusOK).JSON(responseApplRequest)
}

func DeleteApplication(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var application ApplicationM.ApplicantFormRequest

	database.Database.Db.First(&application, id)
	if application.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found with ID"})
	}
	database.Database.Db.Delete(&application)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Application successfully deleted"})
}
