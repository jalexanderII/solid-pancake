package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	c "github.com/jalexanderII/solid-pancake/clients/application/client"
	ApplicationM "github.com/jalexanderII/solid-pancake/clients/application/models"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	commonpb "github.com/jalexanderII/solid-pancake/gen/common"
)

var client = c.NewApplClient()

func Apply(c *fiber.Ctx) error {
	var application *applicationpb.ApplicationReq
	if err := c.BodyParser(&application); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error_message": err.Error()})
	}
	responseApplResponse, err :=  client.Apply(context.Background(), application)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error with review process", "data": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
}

// Upload an attachment
func Upload(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	file, err := c.FormFile("attachment")

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": [1]string{"Unable to upload your attachment"}})
	}
	err = c.SaveFile(file, fmt.Sprintf("./services/application/store/upload/%s", file.Filename))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": [1]string{"Problem saving file"}})
	}

	var appResponse ApplicationM.ApplicantFormResponse
	database.Database.Db.First(&appResponse, id)
	attachments := append(appResponse.Attachments, file.Filename)
	database.Database.Db.Model(&appResponse).Update("attachments", attachments)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Attachment %s uploaded successfully", file.Filename)})
}

func GetApplications(c *fiber.Ctx) error {
	responseApplRequests, err := client.ListApplicationRequests(context.Background(), &commonpb.Empty{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Could not fetch application requests")
	}
	return c.Status(fiber.StatusOK).JSON(responseApplRequests)
}

func GetApplication(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	responseApplRequest, err := client.ReadApplicationRequest(context.Background(), &commonpb.ID{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found for user with that ID"})
	}
	return c.Status(fiber.StatusOK).JSON(responseApplRequest)
}

func DeleteApplication(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	deletedApplication, err := client.DeleteApplicationRequest(context.Background(), &commonpb.ID{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found with ID"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"application": deletedApplication, "message": "Application successfully deleted"})
}

func GetApplicationResponse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	responseApplResponse, err := client.ReadApplicationResponse(context.Background(), &commonpb.ID{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found for user with that ID"})
	}
	return c.Status(fiber.StatusOK).JSON(responseApplResponse)
}

func DeleteApplicationResponse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	deletedApplication, err := client.DeleteApplicationResponse(context.Background(), &commonpb.ID{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No application found with ID"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"application": deletedApplication, "message": "Application successfully deleted"})
}