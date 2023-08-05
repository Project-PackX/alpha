package controllers

import (
	"PackX/initializers"
	"PackX/models"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// Teszt jelleggel az index.html megjelenítése
func PostsIndex(c *fiber.Ctx) error {
	return c.Render("packs/index", fiber.Map{})
}

// Az összes nem törölt DB-ben tárolt csomag adatainak listázása json formában
func ListPackages(c *fiber.Ctx) error {

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
func AddNewPackage(c *fiber.Ctx) error {

	// Csomag előállítása a kérés 'body'-jából
	csomag := new(models.Package)
	if err := c.BodyParser(csomag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// Csomag beszúrása
	result := initializers.DB.Create(&csomag)

	// Hibakezelés
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Hiba történt a csomag létrehozása közben",
		})
	}

	// Alapesetben az új csomag sikeresen létrejött
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Csomag sikeresen hozzáadva",
	})
}

// Csomag törlése
func DeletePackageByID(c *fiber.Ctx) error {

	// ID alapú törlés miatt kell
	// FONTOS: a küldött JSON-ben szám típusú legyen az 'id'
	type DeleteRequest struct {
		ID uint `json:"id"`
	}

	// Csomag előállítása a kérés 'body'-jából
	csomag := new(DeleteRequest)
	if err := c.BodyParser(csomag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// A megadott ID-jú csomag törlése a 'Package'-eket tartalmazó táblából
	result := initializers.DB.Delete(&models.Package{}, csomag.ID)

	// Hibakezelés
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Hiba történt a csomag törlése közben",
		})
	}

	// Alapesetben a csomag sikeresen törlődött
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Csomag sikeresen törölve",
	})
}

// Adott 'id'-jú csomag adatainak lekérdezése
func ListPackageByID(c *fiber.Ctx) error {

	// Az 'id' kinyerése az URL-ből
	id := c.Params("id")

	// Az 'id'-nak megfelelő csomag megkeresése adatbázisból
	var packageData *models.Package
	err := initializers.DB.Where("id = ?", id).First(&packageData).Error

	// Hibakezelés
	if err != nil {
		// Üres a lekérdezés eredménye
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "Nincs ilyen azonosítójú csomag",
			})
		} else { // Minden egyéb..
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Message": "Hiba történt a lekérdezés közben",
			})
		}
		return err
	}

	// A keresett csomag adatainak küldése a kliensnek JSON formában
	c.Status(fiber.StatusOK).JSON(packageData)
	return nil
}
