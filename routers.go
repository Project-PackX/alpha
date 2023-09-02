package main

import (
	"PackX/controllers"

	"github.com/gofiber/fiber/v2"
)

// Setting up endpoints + handle functions
func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex) // Test HTML

	api := app.Group("/api")

	csomagok := api.Group("/packages")
	csomagok.Get("/listuwp", controllers.ListUsersWithPackages)  // /api/csomag/listuwp : Listing packages by users
	csomagok.Get("/list", controllers.ListPackages)              // /api/csomag/list : Listing all packages
	csomagok.Post("/add", controllers.AddNewPackage)             // /api/csomag/add : Inserting new package via input json
	csomagok.Post("/remove", controllers.DeletePackageByID)      // /api/csomag/remove : Delete package based on input json 'id'
	csomagok.Get("/list/:id", controllers.ListPackageByID)       // /api/csomag/list/{id} : Getting the {id}. package details
	csomagok.Get("/getstatus/:id", controllers.GetPackageStatus) // /api/csomag/getstatus/{id} : Getting the {id}. package status

	users := api.Group("/users")
	users.Post("/register", controllers.RegisterNewUser) // /api/users/register : Register new user

	lockers := api.Group("/lockers")
	lockers.Get("/getgroup/:id", controllers.GetCityByLockerID)  // /api/lockers/getgroup/{id} : Get the name of the city where the locker is located at
	lockers.Get("/lgl/:groupid", controllers.ListLockersByGroup) // /api/lockers/lgl/{groupid} : List all the lockers in the {groupid} group
}
