// package handlers
//
// import (
// 	"fmt"
//
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jalexanderII/solid-pancake/database"
// 	"github.com/jalexanderII/solid-pancake/middleware"
// 	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
// 	LifeCycleH "github.com/jalexanderII/solid-pancake/services/lifecycle/handlers"
// )
//
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
//
// // Upload an attachment
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
//
// func CreateFormRequest(application ApplicationM.ApplicantFormRequest) (ApplicantFormRequest, error) {
// 	responseApplRequest := CreateApplicantFormRequest(application)
//
// 	if err := database.Database.Db.Create(&application).Error; err != nil {
// 		return ApplicantFormRequest{}, err
// 	}
// 	responseApplRequest.ID = application.ID
// 	return responseApplRequest, nil
// }
//
// func CreateFormResponse(id int, status string) (ApplicantFormResponse, error) {
// 	var appModel ApplicationM.ApplicantFormRequest
// 	if err := findApplication(id, &appModel); err != nil {
// 		return ApplicantFormResponse{}, err
// 	}
// 	appResponse := ApplicationM.ApplicantFormResponse{
// 		ReferenceId:    uuid.New(),
// 		Status:         status,
// 		ApplicationRef: int(appModel.ID),
// 		Application:    appModel,
// 	}
// 	responseApplResponse := CreateApplicantFormResponse(appResponse)
//
// 	if err := database.Database.Db.Create(&appResponse).Error; err != nil {
// 		return ApplicantFormResponse{}, err
// 	}
// 	responseApplResponse.ID = appResponse.ID
//
// 	applicationRentalDetails := &LifeCycleH.RentalDetailsData{
// 		Data: &LifeCycleH.ApplicationData{ApplicationData: &appResponse},
// 	}
// 	err := LifeCycleH.SendToDataPipeline(applicationRentalDetails)
// 	if err != nil {
// 		return ApplicantFormResponse{}, fmt.Errorf("error processing rental details %v", err)
// 	}
//
// 	return responseApplResponse, nil
// }
//
// func ApplicationReviewProcess(responseApplRequest ApplicantFormRequest) (ApplicantFormResponse, error) {
// 	// TODO: Pass application to review sub-services
// 	status := "PENDING"
// 	return CreateFormResponse(int(responseApplRequest.ID), status)
// }
