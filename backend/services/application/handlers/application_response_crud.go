// package handlers
//
// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jalexanderII/solid-pancake/database"
// 	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
// )
//
// type ApplicantFormResponse struct {
// 	ID          uint                 `json:"id"`
// 	ReferenceId uuid.UUID            `json:"reference_id"`
// 	Status      string               `json:"status,omitempty"`
// 	Attachments []string             `json:"attachments,omitempty"`
// 	Application ApplicantFormRequest `json:"application"`
// }
//
// func CreateApplicantFormResponse(applicantResponseModel ApplicationM.ApplicantFormResponse) ApplicantFormResponse {
// 	var application ApplicationM.ApplicantFormRequest
// 	database.Database.Db.First(&application, applicantResponseModel.ApplicationRef)
// 	return ApplicantFormResponse{
// 		ID:          applicantResponseModel.ID,
// 		ReferenceId: applicantResponseModel.ReferenceId,
// 		Status:      applicantResponseModel.Status,
// 		Attachments: applicantResponseModel.Attachments,
// 		Application: CreateApplicantFormRequest(application),
// 	}
// }
//
// func GetApplicationResponse(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
// 	}
//
// 	var appResponse ApplicationM.ApplicantFormResponse
// 	database.Database.Db.First(&appResponse, id)
// 	if appResponse.ReferenceId == uuid.Nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application response found with ID"})
// 	}
//
// 	responseApplResponse := CreateApplicantFormResponse(appResponse)
//
// 	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
// }
//
// func DeleteApplicationResponse(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
// 	}
//
// 	var appResponse ApplicationM.ApplicantFormResponse
//
// 	database.Database.Db.First(&appResponse, id)
// 	if appResponse.ReferenceId == uuid.Nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application response found with ID"})
// 	}
// 	database.Database.Db.Delete(&appResponse)
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Application Response successfully deleted"})
// }
