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

// Az összes TestDB-ben tárolt adat listázása json formában
func ListItems(c *fiber.Ctx) error {
	// Válasz küldése
	return c.JSON(fiber.Map{
		"message":  "Az elemek sikeresen lekérdezve.",
		"csomagok": initializers.DB,
	})
}

// Csomag hozzáadás
func AddItem(c *fiber.Ctx) error {

	// Statikus csomag készítés a model alapján
	csomag := models.Package{Sender: "Random Pista", Price: 10, Delivered: false}

	// Csomag bezúrása
	result := initializers.DB.Create(&csomag)

	return result.Error
}

// Csomag törlése
func DeleteItem(c *fiber.Ctx) error {

	// Statikus csomag modelje, amit törölni akarunk
	csomag := models.Package{Sender: "Random Pista", Price: 10, Delivered: false}

	// Csomag törlése, aminek id-ja 1
	result := initializers.DB.Delete(&csomag, 1)

	return result.Error
}
