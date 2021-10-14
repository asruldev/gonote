package controllers

import (
	"gonote/app/models"
	"gonote/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	db := database.DB
	body := new(models.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		ID:       uuid.New(),
		Name:     body.Name,
		Email:    body.Email,
		Password: string(password),
	}
	// Create the user and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func LoginUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	return c.JSON(user)
}
