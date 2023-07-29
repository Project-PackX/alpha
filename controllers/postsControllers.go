package controllers

import (
	"PackX/initializers"

	"github.com/gofiber/fiber/v2"
)

// Teszt jelleggel az index.html megjelenítése
func PostsIndex(c *fiber.Ctx) error {
	return c.Render("packs/index", fiber.Map{})
}

var TestDB = initializers.TestDB // A teszt adatbázis használata

// Az összes TestDB-ben tárolt adat listázása json formában
func ListItems(c *fiber.Ctx) error {
	// Válasz küldése
	return c.JSON(fiber.Map{
		"message":  "Az elemek sikeresen lekérdezve.",
		"csomagok": TestDB,
	})
}

// HTTP POST kéréssel json-ben adatátadással új csomag létrehozása
func AddItem(c *fiber.Ctx) error {
	// A kérés testéből kinyerjük az új elemet
	newItem := new(initializers.TestPackage)
	if err := c.BodyParser(newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Hibás kérés",
		})
	}

	// Az új elem hozzáadása az 'adatbázishoz'
	TestDB = append(TestDB, *newItem)

	// Válasz küldése
	return c.JSON(fiber.Map{
		"message": "Az elem sikeresen hozzáadva a listához.",
		"list":    TestDB,
	})
}

// HTTP DELETE kéréssel a megfelelő id-jű elem törlése
func DeleteItem(c *fiber.Ctx) error {

	// A kérés testéből kinyerjük a törlendő elem azonosítóját
	type DeleteRequest struct {
		ID string `json:"id"`
	}

	deleteRequest := new(DeleteRequest)
	if err := c.BodyParser(deleteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Hibás kérés",
		})
	}

	// A törlendő elem keresése a listában
	index := -1
	for i, item := range TestDB {
		if item.ID == deleteRequest.ID {
			index = i
			break
		}
	}

	// Az elem törlése, ha megtalálható
	// A törlés úgy történik, hogy a listát a megfelelő helyen kettészedjük,
	// és újra egyesítjük az 'index'-edik helyen levő elem kihagyásával
	if index >= 0 {
		TestDB = append(TestDB[:index], TestDB[index+1:]...)
	}

	// Válasz küldése
	return c.JSON(fiber.Map{
		"message": "Az elem sikeresen törölve lett.",
		"list":    TestDB,
	})
}
