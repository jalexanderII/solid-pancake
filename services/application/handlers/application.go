package handlers

import (
	"errors"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
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

type ApplicantFormResponse struct {
	ID          uint                 `json:"id"`
	ReferenceId uint32               `json:"reference_id"`
	Status      string               `json:"status,omitempty"`
	Application ApplicantFormRequest `json:"application"`
}

func CreateApplicantFormResponse(applicantResponseModel ApplicationM.ApplicantFormResponse) ApplicantFormResponse {
	var application ApplicationM.ApplicantFormRequest
	database.Database.Db.First(&application, applicantResponseModel.ApplicationRef)
	return ApplicantFormResponse{
		ID:          applicantResponseModel.ID,
		ReferenceId: applicantResponseModel.ReferenceId,
		Status:      applicantResponseModel.Status,
		Application: CreateApplicantFormRequest(application),
	}
}

func Apply(c *fiber.Ctx) error {
	var application ApplicationM.ApplicantFormRequest
	if err := c.BodyParser(&application); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}
	responseApplRequest := CreateApplicantFormRequest(application)
	errs := middleware.ValidateStruct(&responseApplRequest)
	if errs != nil {
		return c.JSON(errs)
	}

	responseApplRequest, err := CreateFormRequest(application)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create application", "data": err.Error()})
	}

	responseApplResponse, err := ApplicationReviewProcess(responseApplRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error with review process", "data": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
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

func CreateFormRequest(application ApplicationM.ApplicantFormRequest) (ApplicantFormRequest, error) {
	responseApplRequest := CreateApplicantFormRequest(application)

	if err := database.Database.Db.Create(&application).Error; err != nil {
		return ApplicantFormRequest{}, err
	}
	responseApplRequest.ID = application.ID
	return responseApplRequest, nil
}

func findApplication(id int, application *ApplicationM.ApplicantFormRequest) error {
	database.Database.Db.Find(&application, "id = ?", id)
	if application.ID == 0 {
		return errors.New("application does not exist")
	}
	return nil
}

func CreateFormResponse(id int, status string) (ApplicantFormResponse, error) {
	var appModel ApplicationM.ApplicantFormRequest
	if err := findApplication(id, &appModel); err != nil {
		return ApplicantFormResponse{}, err
	}
	appResponse := ApplicationM.ApplicantFormResponse{
		ReferenceId:    mutate(appModel.ID),
		Status:         status,
		ApplicationRef: int(appModel.ID),
		Application:    appModel,
	}
	responseApplResponse := CreateApplicantFormResponse(appResponse)

	if err := database.Database.Db.Create(&appResponse).Error; err != nil {
		return ApplicantFormResponse{}, err
	}
	responseApplResponse.ID = appResponse.ID
	return responseApplResponse, nil
}

func mutate(i uint) uint32 {
	return uint32(rand.Intn(100) + int(i)*5)
}

func ApplicationReviewProcess(responseApplRequest ApplicantFormRequest) (ApplicantFormResponse, error) {
	// TODO: Pass application to review sub-services
	status := "PENDING"
	return CreateFormResponse(int(responseApplRequest.ID), status)
}
