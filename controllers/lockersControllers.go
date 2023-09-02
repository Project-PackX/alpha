package controllers

import (
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
)

func GetCityByLockerID(c *fiber.Ctx) error {

	// Get the {id}. row in Locker table
	var locker models.Locker
	initializers.DB.Find(&locker, "id = ?", c.Params("id"))

	// Get the right LockerGroup row
	var lgroup models.LockerGroup
	initializers.DB.Find(&lgroup, "id = ?", locker.LockerGroupID)

	// Sending back the City
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": lgroup.City,
	})
}

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
