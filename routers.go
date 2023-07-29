package main

import (
	"PackX/controllers"

	"github.com/gofiber/fiber/v2"
)

// Végpontok beállítása + hozzá a kezelő függvény
func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)
	app.Get("/list", controllers.ListItems)
	app.Post("/add", controllers.AddItem)
	app.Post("/remove", controllers.DeleteItem)
}
