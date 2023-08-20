package controllers

import (
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Invalid user",
		})
	}

	// Check whether a user with same email exists in the db
	// If the user already exists, return an error
	if initializers.DB.Where("email = ?", newUser.Email).First(&models.User{}).Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "User already exists with given email",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(hashedPassword)

	// Save the user to the database
	initializers.DB.Create(&newUser)

	// Send a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User added successfully",
	})
}
