package main

import (
	"PackX/controllers"

	"github.com/gofiber/fiber/v2"
)

// Végpontok beállítása + hozzá a kezelő függvény
func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex) // Teszt HTML

	api := app.Group("/api")

	csomagok := api.Group("/package")
	csomagok.Get("/listuwp", controllers.ListUsersWithPackages)  // /api/csomag/listuwp : Felhasználónként listázza a csomagokat
	csomagok.Get("/list", controllers.ListPackages)              // /api/csomag/list : Listázza az összes csomagot
	csomagok.Post("/add", controllers.AddNewPackage)             // /api/csomag/add : Új csomag beszúrása
	csomagok.Post("/remove", controllers.DeletePackageByID)      // /api/csomag/remove : A JSON-ben küldött 'id'-jú csomag törlése
	csomagok.Get("/list/:id", controllers.ListPackageByID)       // /api/csomag/list/{id} : Listázza az 'id'-adik számú csomagot
	csomagok.Get("/getstatus/:id", controllers.GetPackageStatus) // /api/csomag/getstatus/{id} : Visszaadja az adott csomag státuszát

	users := api.Group("/users")
	users.Post("/register", controllers.RegisterNewUser) // /api/users/register : Register new user

	lockers := api.Group("/lockers")
	lockers.Get("/getgroup/:id", controllers.GetCityByLockerID)  // /api/lockers/getgroup/{id} : Get the name of the city where the locker is located at
	lockers.Get("/lgl/:groupid", controllers.ListLockersByGroup) // /api/lockers/lgl/{groupid} : List all the lockers in the {groupid} group
}
