package controllers

import (
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewUser(c *fiber.Ctx) error {

	felh := new(models.User)
	if err := c.BodyParser(felh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// Ellenőrizzük, hogy szerepl-e már az ember
	felh1 := models.User{}

	// If the user already exists, return an error
	if initializers.DB.Where("email = ?", felh.Email).First(&felh1).Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Ezzel az email-lel már létezik felhasználó",
		})
	}

	// Save the user to the database
	initializers.DB.Create(&felh)

	// Send a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Felhasználó sikeresen hozzáadva",
	})
}
