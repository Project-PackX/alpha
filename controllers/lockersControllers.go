package controllers

import (
	"PackX/exceptions"
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
)

// Getting the group city based on the locker id via URL
func GetCityByLockerID(c *fiber.Ctx) error {

	// Get the {id}. row in Locker table
	var locker models.Locker
	err := initializers.DB.Where("id = ?", c.Params("id")).First(&locker).Error

	// Check whether or not there is a locker with the given ID
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(exceptions.CreateBaseException(err.Error()))
	}

	// Get the right LockerGroup row
	var lgroup models.LockerGroup
	initializers.DB.Find(&lgroup, "id = ?", locker.ID[:2])

	// Sending back the name of the city
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": lgroup.City,
	})
}

// Get all the package in the input locker via URL id
func GetPackagesByLockerID(c *fiber.Ctx) error {

	// Getting the {id} from URL
	id := c.Params("id")

	// Getting the connection models which have the right locker id
	var temp []models.PackageLocker
	initializers.DB.Find(&temp, "locker_id = ?", id)

	// getting a slice of package ids, which are in the locker
	var packages []uint
	for i := 0; i < len(temp); i++ {
		packages = append(packages, temp[i].Package_id)
	}

	// Sending back the list of package ids
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": packages,
	})
}
