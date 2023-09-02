package controllers

import (
	"PackX/exceptions"
	"PackX/initializers"
	"PackX/models"

	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(c *fiber.Ctx) error {

	// Creating the new user
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.BaseException{
			Message:   "Invalid user",
			TimeStamp: time.Now().String(),
		})
	}

	// Check whether a user with same email exists in the db
	// If the user already exists, return an error
	if initializers.DB.Where("email = ?", newUser.Email).First(&models.User{}).Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.BaseException{
			Message:   "User already exists with given email",
			TimeStamp: time.Now().String(),
		})
	}

	// Password hashing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(hashedPassword)

	// Save the user to the database
	initializers.DB.Create(&newUser)

	// Send a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User added successfully",
	})
}
