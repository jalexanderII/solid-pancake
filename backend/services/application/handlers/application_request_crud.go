package handlers

import (
	"errors"
	"fmt"

	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	commonpb "github.com/jalexanderII/solid-pancake/gen/common"
	RealEstateH "github.com/jalexanderII/solid-pancake/services/realestate/handlers"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	UserH "github.com/jalexanderII/solid-pancake/services/users/handlers"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
)

type ApplicantFormRequest struct {
	ID              uint                  `json:"id"`
	Name            string                `json:"name"`
	SocialSecurity  string                `json:"social_security,omitempty"`
	DateOfBirth     string                `json:"date_of_birth,omitempty"`
	DriversLicense  string                `json:"drivers_license,omitempty"`
	PreviousAddress *commonpb.Place       `json:"previous_address,omitempty" validate:"dive"`
	Landlord        string                `json:"landlord,omitempty"`
	LandlordNumber  string                `json:"landlord_number,omitempty"`
	Employer        string                `json:"employer,omitempty"`
	Salary          int32                 `json:"salary,omitempty"`
	Apartment       RealEstateH.Apartment `json:"apartment" validate:"dive"`
	User            UserH.User            `json:"user" validate:"dive"`
}

// CreateApplicantFormRequest Takes in a model and returns a serializer
func CreateApplicantFormRequest(applicantRequestModel *applicationpb.ApplicationReq) *applicationpb.ApplicationReq {
	var (
		apartment RealEstateM.Apartment
		user      UserM.User
	)
	database.Database.Db.First(&apartment, applicantRequestModel.ApartmentRef)
	database.Database.Db.First(&user, applicantRequestModel.UserRef)
	return &applicationpb.ApplicationReq{
		Id:              applicantRequestModel.Id,
		Name:            applicantRequestModel.Name,
		UserRef:         int32(user.ID),
		SocialSecurity:  applicantRequestModel.SocialSecurity,
		DateOfBirth:     applicantRequestModel.DateOfBirth,
		DriversLicense:  applicantRequestModel.DriversLicense,
		PreviousAddress: applicantRequestModel.PreviousAddress,
		Landlord:        applicantRequestModel.Landlord,
		LandlordNumber:  applicantRequestModel.LandlordNumber,
		Employer:        applicantRequestModel.Employer,
		Salary:          applicantRequestModel.Salary,
		ApartmentRef:    int32(apartment.ID),
	}
}

func (h *Handler) GetApplications() (*applicationpb.ListApplicationReqOut, error) {
	var applications []*applicationpb.ApplicationReq
	database.Database.Db.Find(&applications)

	responseApplRequests := make([]*applicationpb.ApplicationReq, len(applications))
	for idx, application := range applications {
		responseApplRequests[idx] = CreateApplicantFormRequest(application)
	}
	return &applicationpb.ListApplicationReqOut{ApplicationRequests: responseApplRequests}, nil
}

// func GetApplications(c *fiber.Ctx) error {
// 	var applications []*applicationpb.ApplicationReq
// 	database.Database.Db.Find(&applications)
//
// 	responseApplRequests := make([]*applicationpb.ApplicationReq, len(applications))
// 	for idx, application := range applications {
// 		responseApplRequests[idx] = CreateApplicantFormRequest(application)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(responseApplRequests)
// }

func findApplication(id int, application *applicationpb.ApplicationRes) error {
	database.Database.Db.Find(&application, "id = ?", id)
	if application.Id == 0 {
		return errors.New("application does not exist")
	}
	return nil
}

func (h *Handler) GetApplication(id int32) (*applicationpb.ApplicationReq, error) {
	var application *applicationpb.ApplicationReq
	database.Database.Db.Find(&application, "user_id = ?", id)
	if application.Id == 0 {
		return nil, fmt.Errorf("no application found four user with that ID")
	}

	responseApplRequest := CreateApplicantFormRequest(application)

	return responseApplRequest, nil
}
//
// func GetApplication(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
// 	}
//
// 	var application *applicationpb.ApplicationReq
// 	database.Database.Db.Find(&application, "user_id = ?", id)
// 	if application.Id == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found four user with that ID"})
// 	}
//
// 	responseApplRequest := CreateApplicantFormRequest(application)
//
// 	return c.Status(fiber.StatusOK).JSON(responseApplRequest)
// }

func (h *Handler) DeleteApplication(id int32) (*applicationpb.ApplicationReq, error) {
	var application *applicationpb.ApplicationReq

	database.Database.Db.First(&application, id)
	if application.Name == "" {
		return nil, fmt.Errorf("no application found with ID")
	}
	database.Database.Db.Delete(&application)
	return application, nil
}
//
// func DeleteApplication(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
// 	}
//
// 	var application ApplicationM.ApplicantFormRequest
//
// 	database.Database.Db.First(&application, id)
// 	if application.Name == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found with ID"})
// 	}
// 	database.Database.Db.Delete(&application)
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Application successfully deleted"})
// }
