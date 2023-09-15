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

	csomagok.Get("", controllers.ListPackages)                   // /api/packages : Listing all packages
	csomagok.Post("", controllers.AddNewPackage)                 // /api/packages : Inserting new package via input json
	csomagok.Delete("/:id", controllers.DeletePackageByID)       // /api/packages/{id} : Delete package based on pathvariable 'id'
	csomagok.Get("/:id", controllers.ListPackageByID)            // /api/packages/{id} : Getting the {id}. package details
	csomagok.Get("/getstatus/:id", controllers.GetPackageStatus) // /api/packages/getstatus/{id} : Getting the {id}. package status
	csomagok.Post("/change-status", controllers.ChangeStatus)    // /api/packages/change-status : Change a package status via input JSON (ID, NewStatusID)
	csomagok.Post("/cancel/:id", controllers.MakeCanceled)       // /api/packages/cancel/{id} : Make canceled a package based on pathvariable 'id'

	users := api.Group("/users")
	users.Post("/register", controllers.RegisterNewUser) // /api/users/register : Register new user via input JSON
	users.Post("/login", controllers.Login)
	// users.Get("/packages", controllers.GetPackagesUnderUsers) "// csomagok.Get("/uwp", controllers.ListUsersWithPackages" Instead of this, use the users/packages or just get all of the packages
	// if you want to get a specific user's packages, refer to the next one
	// users.Get("/:id/packages", controller.GetPackagesUnderUser) // /api/users/{id}/packages : Get all packages under user

	lockers := api.Group("/lockers")
	lockers.Get("/get-city/:id", controllers.GetCityByLockerID) // /api/lockers/{id} : Get the name of the city where the locker is located at
	lockers.Get("/get-packages/:id", controllers.GetPackagesByLockerID)

	/* I guess we don't need this anymore, but I leave it here
	lockers.Get("/lockers-by-group/:groupid", controllers.ListLockersByGroup) // /api/lockers/lockers-by-group/{groupid} : List all the lockers in the {groupid} group
	*/
}
