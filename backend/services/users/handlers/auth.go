package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/config"
	"github.com/jalexanderII/solid-pancake/database"
	"github.com/jalexanderII/solid-pancake/middleware"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *fiber.Ctx) error {
	var user UserM.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	user.Password = hash

	responseUser := CreateResponseUser(user)
	errs := middleware.ValidateStruct(&responseUser)
	if errs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": errs})
	}

	if err := database.Database.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err.Error()})
	}
	responseUser.ID = user.ID

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	email := input.Email
	pass := input.Password

	user, err := getUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if !CheckPasswordHash(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token, err := config.GenerateNewAccessToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not login"})
	}

	cookie := config.GenerateNewCookie(token, true)
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": token})
}

func Logout(c *fiber.Ctx) error {
	cookie := config.GenerateNewCookie("", false)
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"status": "success", "message": "Success logout"})
}

func authUser(c *fiber.Ctx) (UserM.User, error) {
	var user UserM.User
	claimIssuerID, err := config.CheckToken(c)
	if err != nil {
		return user, c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err = findUser(int(claimIssuerID), &user); err != nil {
		return user, c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return user, nil
}

func validUser(id int, p string) bool {
	var user UserM.User
	database.Database.Db.First(&user, id)
	if user.Password == "" {
		return true
	}
	if user.Username == "" {
		return false
	}

	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}
