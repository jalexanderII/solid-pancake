package handlers

import (
	"errors"
	"fmt"

	ApplicationM "github.com/jalexanderII/solid-pancake/clients/application/models"
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
func CreateApplicantFormRequest(applicantRequestModel ApplicationM.ApplicantFormRequest) *applicationpb.ApplicationReq {
	var (
		apartment RealEstateM.Apartment
		user      UserM.User
	)
	database.Database.Db.First(&apartment, applicantRequestModel.ApartmentRef)
	database.Database.Db.First(&user, applicantRequestModel.UserRef)
	return &applicationpb.ApplicationReq{
		Id:             int32(applicantRequestModel.ID),
		Name:           applicantRequestModel.Name,
		UserRef:        int32(user.ID),
		SocialSecurity: applicantRequestModel.SocialSecurity,
		DateOfBirth:    applicantRequestModel.DateOfBirth,
		DriversLicense: applicantRequestModel.DriversLicense,
		PreviousAddress: &commonpb.Place{
			Address:      applicantRequestModel.PreviousAddress.Address,
			Street:       applicantRequestModel.PreviousAddress.Street,
			City:         applicantRequestModel.PreviousAddress.City,
			State:        applicantRequestModel.PreviousAddress.State,
			Zip:          applicantRequestModel.PreviousAddress.Zip,
			Neighborhood: applicantRequestModel.PreviousAddress.Neighborhood,
			Unit:         applicantRequestModel.PreviousAddress.Unit,
			Lat:          applicantRequestModel.PreviousAddress.Lat,
			Lng:          applicantRequestModel.PreviousAddress.Lng,
		},
		Landlord:       applicantRequestModel.Landlord,
		LandlordNumber: applicantRequestModel.LandlordNumber,
		Employer:       applicantRequestModel.Employer,
		Salary:         applicantRequestModel.Salary,
		ApartmentRef:   int32(apartment.ID),
	}
}

func (h *Handler) GetApplications() (*applicationpb.ListApplicationReqOut, error) {
	var applications []ApplicationM.ApplicantFormRequest
	database.Database.Db.Find(&applications)

	responseApplRequests := make([]*applicationpb.ApplicationReq, len(applications))
	for idx, application := range applications {
		responseApplRequests[idx] = CreateApplicantFormRequest(application)
	}
	return &applicationpb.ListApplicationReqOut{ApplicationRequests: responseApplRequests}, nil
}

func FindApplication(id int, application ApplicationM.ApplicantFormRequest) error {
	database.Database.Db.Find(&application, "id = ?", id)
	if application.ID == 0 {
		return errors.New("application does not exist")
	}
	return nil
}

func (h *Handler) GetApplication(id int32) (*applicationpb.ApplicationReq, error) {
	var application ApplicationM.ApplicantFormRequest
	database.Database.Db.Find(&application, "user_id = ?", id)
	if application.ID == 0 {
		return nil, fmt.Errorf("no application found four user with that ID")
	}

	responseApplRequest := CreateApplicantFormRequest(application)

	return responseApplRequest, nil
}

func (h *Handler) DeleteApplication(id int32) (*applicationpb.ApplicationReq, error) {
	var application ApplicationM.ApplicantFormRequest

	database.Database.Db.First(&application, id)
	if application.Name == "" {
		return nil, fmt.Errorf("no application found with ID")
	}
	database.Database.Db.Delete(&application)
	return CreateApplicantFormRequest(application), nil
}
