package controllers

import (
	"math"

	"github.com/Project-PackX/backend/exceptions"
	"github.com/Project-PackX/backend/initializers"
	"github.com/Project-PackX/backend/models"

	"github.com/gofiber/fiber/v2"
)

func AddNewLocker(c *fiber.Ctx) error {
	// Creating the new locker with the input json body
	locker := new(models.Locker)
	if err := c.BodyParser(locker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad request",
		})
	}

	// Inserting the new package
	result := initializers.DB.Create(&locker)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "There was an error during adding new locker",
		})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Locker successfully added",
	})
}

// List all lockers
func ListLockers(c *fiber.Ctx) error {

	var lockers []models.Locker // Slice that will contain all lockers

	// Execute 'SELECT * FROM public.lockers' query
	result := initializers.DB.Find(&lockers)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Something bad happened during the query",
		})
	}

	// Separately treated case where no record exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "There are no lockers in the database",
		})
	}

	var npacks []int
	var percents []float64

	for i := 0; i < len(lockers); i++ {

		var temp []models.PackageLocker
		initializers.DB.Find(&temp, "locker_id = ?", lockers[i].ID)

		nPackages := int(len(temp))

		percent := float64(nPackages) / float64(lockers[i].Capacity)
		percent = math.Round(percent * 100)

		npacks = append(npacks, nPackages)
		percents = append(percents, float64(percent))

	}

	// Returning the lockers
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "Success",
		"lockers":          lockers,
		"numberofpackages": npacks,
		"percents":         percents,
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
	var packagesToLockers []models.PackageLocker
	initializers.DB.Find(&packagesToLockers, "locker_id = ?", id)

	nPackages := len(packagesToLockers)

	percent := float64(nPackages) / float64(cap)
	percent = math.Round(percent * 100)

	// Return the datas
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Capacity":      cap,
		"PackageNumber": nPackages,
		"Percent":       percent,
	})
}
