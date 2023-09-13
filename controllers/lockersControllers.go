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

// List all lockers by the input group id
func ListLockersByGroup(c *fiber.Ctx) error {

	// Get the locker group id from the request
	id := c.Params("groupid")

	// Getting all lockers with the right group id
	var lockers []models.Locker
	initializers.DB.Find(&lockers, "locker_group_id = ?", id)

	// Sending back all the lockers in that group with JSON format
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": lockers,
	})
}
