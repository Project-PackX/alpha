package main

import (
	"PackX/controllers"
	"PackX/middleware"

	"github.com/gofiber/fiber/v2"
)

// Setting up endpoints + handle functions
func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex) // Test HTML

	api := app.Group("/api")

	packages := api.Group("/packages")

	packages.Get("/get/:trackid", controllers.ListPackageByID) // /api/packages/get/{id} : Getting the {id}. package details

	// From this point, all package endpoints are being authenticated
	packages.Use(middleware.RequireJwtTokenAuth)

	packages.Get("/all", controllers.ListPackages)               // /api/packages/all : Listing all packages
	packages.Post("/new", controllers.AddNewPackage)             // /api/packages/new : Inserting new package via input json
	packages.Delete("/:id", controllers.DeletePackageByID)       // /api/packages/{id} : Delete package based on pathvariable 'id'
	packages.Get("/getstatus/:id", controllers.GetPackageStatus) // /api/packages/getstatus/{id} : Getting the {id}. package status
	packages.Post("/change-status", controllers.ChangeStatus)    // /api/packages/change-status : Change a package status via input JSON (ID, NewStatusID)
	packages.Post("/cancel/:id", controllers.MakeCanceled)       // /api/packages/cancel/{id} : Make canceled a package based on pathvariable 'id'

	users := api.Group("/users")

	users.Get("/all", controllers.ListUsers)             // /api/users/all : Listing all users
	users.Post("/register", controllers.RegisterNewUser) // /api/users/register : Register new user via input JSON

	users.Post("/login", controllers.Login) // /api/users/login : Login user

	// From this point, all user endpoints are being authenticated
	users.Use(middleware.RequireJwtTokenAuth)

	users.Get("/get-accesslevel/:id", controllers.GetAccessLevel) // /api/users/get-accesslevel/{id} : Get the access level of the {id}. user
	users.Post("/set-accesslevel", controllers.SetAccessLevel)    // /api/users/set-accesslevel : Set the access level of the user {email, accesslevel}
	/*
		Until we find a better approach for this accesslevel problem...
		Should we send a number, or string, maybe create a new DB table with these pairs?
	*/
	users.Get("/:id/packages", controllers.GetPackagesUnderUser) // /api/users/{id}/packages : Get all packages which the {id}. user sent

	// users.Get("/packages", controllers.GetPackagesUnderUsers) "// csomagok.Get("/uwp", controllers.ListUsersWithPackages" Instead of this, use the users/packages or just get all of the packages

	lockers := api.Group("/lockers")

	lockers.Get("/all", controllers.ListLockers) // api/lockers/all : Listing all lockers

	// From this point, all locker endpoints are being authenticated
	lockers.Use(middleware.RequireJwtTokenAuth)

	lockers.Get("/get-packages/:id", controllers.GetPackagesByLockerID) // /api/lockers/get-packages/{id} : Get all the information about the packages that are in the {id}. locker
	lockers.Get("/get-fullness/:id", controllers.GetFullness)           // /api/lockers/get-fullness/{id} : Get the fullness stats (cap, number of package, percentage) of the {id}. locker

	/* I guess we don't need this anymore, but I leave it here
	lockers.Get("/lockers-by-group/:groupid", controllers.ListLockersByGroup) // /api/lockers/lockers-by-group/{groupid} : List all the lockers in the {groupid} group
	*/
}
