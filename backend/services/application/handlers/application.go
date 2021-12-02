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

// func Apply(c *fiber.Ctx) error {
// 	var application ApplicationM.ApplicantFormRequest
// 	if err := c.BodyParser(&application); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
// 	}
// 	responseApplRequest := CreateApplicantFormRequest(application)
// 	errs := middleware.ValidateStruct(&responseApplRequest)
// 	if errs != nil {
// 		return c.JSON(errs)
// 	}
//
// 	responseApplRequest, err := CreateFormRequest(application)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create application", "data": err.Error()})
// 	}
//
// 	responseApplResponse, err := ApplicationReviewProcess(responseApplRequest)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error with review process", "data": err.Error()})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
// }

// Upload an attachment
// func Upload(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
// 	}
//
// 	file, err := c.FormFile("attachment")
//
// 	if err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": [1]string{"Unable to upload your attachment"}})
// 	}
// 	err = c.SaveFile(file, fmt.Sprintf("./services/application/store/upload/%s", file.Filename))
// 	if err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": [1]string{"Problem saving file"}})
// 	}
//
// 	var appResponse ApplicationM.ApplicantFormResponse
// 	database.Database.Db.First(&appResponse, id)
// 	attachments := append(appResponse.Attachments, file.Filename)
// 	database.Database.Db.Model(&appResponse).Update("attachments", attachments)
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Attachment %s uploaded successfully", file.Filename)})
// }

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
