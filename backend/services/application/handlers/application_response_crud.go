package handlers

import (
	"fmt"

	"github.com/google/uuid"
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

func CreateApplicantFormResponse(applicantResponseModel *applicationpb.ApplicationRes) *applicationpb.ApplicationRes {
	var application *applicationpb.ApplicationReq
	database.Database.Db.First(&application, applicantResponseModel.ApplicationRef)
	return &applicationpb.ApplicationRes {
		Id:             applicantResponseModel.Id,
		ReferenceId:    &common.UUID{Value: applicantResponseModel.ReferenceId.String()},
		Status:         applicantResponseModel.Status,
		Attachments:    applicantResponseModel.Attachments,
		ApplicationRef: application.Id,
	}
}

func (h *Handler) GetApplicationResponse(id int32) (*applicationpb.ApplicationRes, error) {
	var appResponse *applicationpb.ApplicationRes
	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId.String() == "" {
		return nil, fmt.Errorf("no application response found with ID")
	}
	responseApplResponse := CreateApplicantFormResponse(appResponse)

	return responseApplResponse, nil
}

func (h *Handler) DeleteApplicationResponse(id int32) (*applicationpb.ApplicationRes, error) {
	var appResponse *applicationpb.ApplicationRes

	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId.String() == "" {
		return nil, fmt.Errorf("no application response found with ID")
	}
	database.Database.Db.Delete(&appResponse)
	return appResponse, nil
}
