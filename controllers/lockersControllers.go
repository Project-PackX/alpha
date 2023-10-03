package controllers

import (
	"PackX/exceptions"
	"PackX/initializers"
	"PackX/models"
	"math"

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
	var ids []uint
	for i := 0; i < len(temp); i++ {
		ids = append(ids, temp[i].Package_id)
	}

	// If there are no packages in the locker
	if ids == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message": "There are no packages in this locker",
		})
	}

	// THe slice of the packages that are in the specific locker
	var packs []models.Package

	// Execute 'SELECT * FROM public.csomagok' query
	result := initializers.DB.Find(&packs, ids)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateBaseException("Something went wrong during query"))
	}

	// Sending back the list of packages with every information
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": packs,
	})
}

// Get the fullness of a locker based on URL
func GetFullness(c *fiber.Ctx) error {

	// Getting the id from URL
	id := c.Params("id")

	// Getting the locker's capacity
	var locker models.Locker
	initializers.DB.Where("id = ?", id).First(&locker)

	cap := locker.Capacity

	// Getting the number of packages in the locker
	var temp []models.PackageLocker
	initializers.DB.Find(&temp, "locker_id = ?", id)

	nPackages := len(temp)

	percent := float64(nPackages) / float64(cap)
	percent = math.Round(percent * 100)

	// Return the datas
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Capacity":      cap,
		"PackageNumber": nPackages,
		"Percent":       percent,
	})
}
