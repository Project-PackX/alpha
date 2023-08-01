package controllers

import (
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
)

// Teszt jelleggel az index.html megjelenítése
func PostsIndex(c *fiber.Ctx) error {
	return c.Render("packs/index", fiber.Map{})
}

// Az összes nem törölt DB-ben tárolt csomag adatainak listázása json formában
func ListItems(c *fiber.Ctx) error {

	var csomagok []models.Package // A lekérdezéshez kell

	// Futtat egy 'SELECT * FROM public.csomagok' lekérdezést
	result := initializers.DB.Find(&csomagok)

	// Hibakezelés
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hiba történt a lekérdezés közben",
		})
	}

	// Külön kezelt eset amikor egy rekord sincsen
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message": "Nincs egy aktív csomag sem",
		})
	}

	// Alapesetben normál válasz a rekordok küldésével
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message":  "Success",
		"Csomagok": csomagok,
	})
}

// Csomag hozzáadása
func AddItem(c *fiber.Ctx) error {

	// Csomag előállítása a kérés 'body'-jából
	csomag := new(models.Package)
	if err := c.BodyParser(csomag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// Csomag bezúrása
	result := initializers.DB.Create(&csomag)

	return result.Error
}

// Csomag törlése
func DeleteItem(c *fiber.Ctx) error {

	// ID alapú törlés miatt kell
	type DeleteRequest struct {
		ID string `json:"id"`
	}

	// Csomag előállítása a kérés 'body'-jából
	csomag := new(DeleteRequest)
	if err := c.BodyParser(csomag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// A megadott ID-ú csomag törlése a 'Package'-eket tartalmazó táblából
	result := initializers.DB.Delete(&models.Package{}, csomag.ID)

	return result.Error
}
