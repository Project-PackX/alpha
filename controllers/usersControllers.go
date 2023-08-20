package controllers

import (
	"PackX/initializers"
	"PackX/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewUser(c *fiber.Ctx) error {
	// Felhasználóhoz szükséges infók
	newName := c.Params("name")
	newEmail := c.Params("email")
	newPassword := c.Params("password")

	fmt.Println(newEmail)

	// Ellenőrizzük, hogy szerepl-e már az ember
	user := models.User{}
	err := initializers.DB.Where("email = ?", newEmail).First(&user).Error

	// If the user already exists, return an error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Ezzel az email-lel már létezik felhasználó",
		})
	}

	// Create a new user
	newUser := models.User{
		Name:     newName,
		Email:    newEmail,
		Password: newPassword,
	}

	// Save the user to the database
	initializers.DB.Create(&newUser)

	// Send a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Felhasználó sikeresen hozzáadva",
	})
}
