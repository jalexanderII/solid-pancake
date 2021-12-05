package handlers

import (
	"fmt"

	"github.com/google/uuid"
	ApplicationM "github.com/jalexanderII/solid-pancake/clients/application/models"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	LifeCycleH "github.com/jalexanderII/solid-pancake/services/lifecycle/handlers"
)

type Handler struct{}

func NewHandler() *Handler {
	h := &Handler{}
	return h
}

func (h *Handler) Apply(application ApplicationM.ApplicantFormRequest) (*applicationpb.ApplicationRes, error) {
	responseApplRequest, err := CreateFormRequest(application)
	if err != nil {
		return nil, err
	}

	responseApplResponse, err := ApplicationReviewProcess(responseApplRequest)
	if err != nil {
		return nil, err
	}

	return responseApplResponse, nil
}

func CreateFormRequest(application ApplicationM.ApplicantFormRequest) (*applicationpb.ApplicationReq, error) {
	responseApplRequest := CreateApplicantFormRequest(application)

	if err := database.Database.Db.Create(&application).Error; err != nil {
		return nil, err
	}
	responseApplRequest.Id = int32(application.ID)
	return responseApplRequest, nil
}

func CreateFormResponse(id int, status string) (*applicationpb.ApplicationRes, error) {
	var appModel ApplicationM.ApplicantFormRequest
	if err := FindApplication(id, appModel); err != nil {
		return nil, err
	}

	appResponse := ApplicationM.ApplicantFormResponse{
		ReferenceId:    uuid.New(),
		Status:         status,
		Attachments:    []string{},
		ApplicationRef: int(appModel.ID),
		Application:    appModel,
	}
	responseApplResponse := CreateApplicantFormResponse(appResponse)

	if err := database.Database.Db.Create(&appResponse).Error; err != nil {
		return nil, err
	}

	responseApplResponse.Id = int32(appResponse.ID)

	applicationRentalDetails := &LifeCycleH.RentalDetailsData{
		Data: &LifeCycleH.ApplicationData{ApplicationData: &appResponse},
	}
	err := LifeCycleH.SendToDataPipeline(applicationRentalDetails)
	if err != nil {
		return nil, fmt.Errorf("error processing rental details %v", err)
	}

	return responseApplResponse, nil
}

func ApplicationReviewProcess(responseApplRequest *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error) {
	// TODO: Pass application to review sub-services
	status := "PENDING"
	return CreateFormResponse(int(responseApplRequest.Id), status)
}
