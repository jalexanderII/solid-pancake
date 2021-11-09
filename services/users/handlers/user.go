package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User To be used as a serializer
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
}

// CreateResponseUser Takes in a model and returns a serializer
func CreateResponseUser(userModel UserM.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Username: userModel.Username, Email: userModel.Email}
}

func getUserByEmail(e string) (*UserM.User, error) {
	var user UserM.User
	if err := database.Database.Db.Where(&UserM.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func findUser(id int, user *UserM.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	var users []UserM.User
	responseUsers := make([]User, len(users))

	database.Database.Db.Find(&users)
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var user UserM.User

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

type UpdateUserResponse struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UpdateUser(c *fiber.Ctx) error {
	var user UserM.User
	var updateUserResponse UpdateUserResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	user, err = authUser(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	// user can only update their own info
	if id != int(user.ID) {
		return c.Status(fiber.StatusBadRequest).JSON("Can't update another user")
	}

	if err = c.BodyParser(&updateUserResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	database.Database.Db.Model(&user).Clauses(clause.Returning{}).Updates(updateUserResponse)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	type DeleteUser struct {
		Password string `json:"password"`
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var deleteUser DeleteUser
	user, err := authUser(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	// user can only update their own info
	if id != int(user.ID) {
		return c.Status(fiber.StatusBadRequest).JSON("Can't delete another user")
	}

	if err := c.BodyParser(&deleteUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	if !validUser(int(user.ID), deleteUser.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}

	if err := database.Database.Db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}
