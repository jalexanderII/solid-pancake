package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/gen/common"
	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
)

type ApplicantFormResponse struct {
	ID          uint                 `json:"id"`
	ReferenceId uuid.UUID            `json:"reference_id"`
	Status      string               `json:"status,omitempty"`
	Attachments []string             `json:"attachments,omitempty"`
	Application ApplicantFormRequest `json:"application"`
}

func CreateApplicantFormResponse(applicantResponseModel ApplicationM.ApplicantFormResponse) *applicationpb.ApplicationRes {
	var application *applicationpb.ApplicationReq
	database.Database.Db.First(&application, applicantResponseModel.ApplicationRef)
	return &applicationpb.ApplicationRes {
		Id: int32(applicantResponseModel.ID),
		ReferenceId:    &common.UUID{Value: applicantResponseModel.ReferenceId.String()},
		Status:         applicantResponseModel.Status,
		Attachments:    applicantResponseModel.Attachments,
		ApplicationRef: application.Id,
	}
}

func GetApplicationResponse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var appResponse ApplicationM.ApplicantFormResponse
	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application response found with ID"})
	}

	responseApplResponse := CreateApplicantFormResponse(appResponse)

	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
}

func DeleteApplicationResponse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var appResponse ApplicationM.ApplicantFormResponse

	database.Database.Db.First(&appResponse, id)
	if appResponse.ReferenceId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application response found with ID"})
	}
	database.Database.Db.Delete(&appResponse)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Application Response successfully deleted"})
}
