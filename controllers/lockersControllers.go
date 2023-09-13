package controllers

import (
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
)

// Handy function
func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

// Getting the group city based on the locker id via URL
func GetCityByLockerID(c *fiber.Ctx) error {

	// Get the {id}. row in Locker table
	var locker models.Locker
	initializers.DB.Find(&locker, "id = ?", c.Params("id"))

	//Getting the lockergroup code
	groupCode := firstN(locker.ID, 2)

	// Get the right LockerGroup row
	var lgroup models.LockerGroup
	initializers.DB.Find(&lgroup, "id = ?", groupCode)

	// Sending back the City
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
