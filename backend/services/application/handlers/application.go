package handlers

import (
	"github.com/google/uuid"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	commonpb "github.com/jalexanderII/solid-pancake/gen/common"
)

type Handler struct {}

func NewHandler() *Handler {
	h := &Handler{}
	return h
}

func (h *Handler) Apply(application *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error) {
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


func CreateFormRequest(application *applicationpb.ApplicationReq) (*applicationpb.ApplicationReq, error) {
	responseApplRequest := CreateApplicantFormRequest(application)

	if err := database.Database.Db.Create(&application).Error; err != nil {
		return &applicationpb.ApplicationReq{}, err
	}
	responseApplRequest.Id = application.Id
	return responseApplRequest, nil
}

func CreateFormResponse(id int, status string) (*applicationpb.ApplicationRes, error) {
	var appModel *applicationpb.ApplicationRes
	if err := findApplication(id, appModel); err != nil {
		return nil, err
	}

	appResponse := &applicationpb.ApplicationRes{
		ReferenceId:    &commonpb.UUID{Value: uuid.NewString()},
		Status:         status,
		Attachments: appModel.Attachments,
		ApplicationRef: appModel.Id,
	}
	responseApplResponse := CreateApplicantFormResponse(appResponse)

	if err := database.Database.Db.Create(&appResponse).Error; err != nil {
		return nil, err
	}

	responseApplResponse.Id = appResponse.Id

	// applicationRentalDetails := &LifeCycleH.RentalDetailsData{
	// 	Data: &LifeCycleH.ApplicationData{ApplicationData: appResponse},
	// }
	// err := LifeCycleH.SendToDataPipeline(applicationRentalDetails)
	// if err != nil {
	// 	return nil, fmt.Error("error processing rental details %v", err)
	// }

	return responseApplResponse, nil
}

func ApplicationReviewProcess(responseApplRequest *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error) {
	// TODO: Pass application to review sub-services
	status := "PENDING"
	return CreateFormResponse(int(responseApplRequest.Id), status)
}
