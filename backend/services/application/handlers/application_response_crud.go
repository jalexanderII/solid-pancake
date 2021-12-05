package handlers

import (
	"fmt"

	"github.com/google/uuid"
	ApplicationM "github.com/jalexanderII/solid-pancake/clients/application/models"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/gen/common"
)

type ApplicantFormResponse struct {
	ID          uint                 `json:"id"`
	ReferenceId uuid.UUID            `json:"reference_id"`
	Status      string               `json:"status,omitempty"`
	Attachments []string             `json:"attachments,omitempty"`
	Application ApplicantFormRequest `json:"application"`
}

func CreateApplicantFormResponse(applicantResponseModel ApplicationM.ApplicantFormResponse) *applicationpb.ApplicationRes {
	var application ApplicationM.ApplicantFormRequest
	database.Database.Db.First(&application, applicantResponseModel.ApplicationRef)
	return &applicationpb.ApplicationRes{
		Id:             int32(applicantResponseModel.ID),
		ReferenceId:    &common.UUID{Value: applicantResponseModel.ReferenceId.String()},
		Status:         applicantResponseModel.Status,
		Attachments:    applicantResponseModel.Attachments,
		ApplicationRef: int32(application.ID),
	}
}

func (h *Handler) GetApplicationResponse(id int32) (*applicationpb.ApplicationRes, error) {
	var appResponse ApplicationM.ApplicantFormResponse
	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId.String() == "" {
		return nil, fmt.Errorf("no application response found with ID")
	}
	responseApplResponse := CreateApplicantFormResponse(appResponse)

	return responseApplResponse, nil
}

func (h *Handler) DeleteApplicationResponse(id int32) (*applicationpb.ApplicationRes, error) {
	var appResponse ApplicationM.ApplicantFormResponse

	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId.String() == "" {
		return nil, fmt.Errorf("no application response found with ID")
	}
	database.Database.Db.Delete(&appResponse)
	return CreateApplicantFormResponse(appResponse), nil
}
