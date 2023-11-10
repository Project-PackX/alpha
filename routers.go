package main

import (
	"github.com/Project-PackX/backend/controllers"
	"github.com/Project-PackX/backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// Setting up endpoints + handle functions
func Routes(app *fiber.App) {
	api := app.Group("/api")
	packages := api.Group("/packages")

	packages.Get("/get/:trackid", controllers.ListPackageByID) // /api/packages/get/{id} : Getting the {id}. package details

	// From this point, all package endpoints are being authenticated
	packages.Use(middleware.RequireJwtTokenAuth)

	packages.Get("/all", controllers.ListPackages)                            // /api/packages/all : Listing all packages
	packages.Post("/new", controllers.AddNewPackage)                          // /api/packages/new : Inserting new package via input json
	packages.Delete("/:id", controllers.DeletePackageByID)                    // /api/packages/{id} : Delete package based on pathvariable 'id'
	packages.Get("/getstatus/:id", controllers.GetPackageStatus)              // /api/packages/getstatus/{id} : Getting the {id}. package status
	packages.Post("/statusup/:id", controllers.ChangeStatusUp)                // /api/packages/statusup/{id} : Increnent {id}. package status
	packages.Get("/courierpackages/:id", controllers.GetPackagesUnderCourier) // /api/packages/courierpackages/:id : Get packages under desired courier
	packages.Post("/change-status", controllers.ChangeStatus)                 // /api/packages/change-status : Change a package status via input JSON (ID, NewStatusID)
	packages.Post("/cancel/:id", controllers.MakeCanceled)                    // /api/packages/cancel/{id} : Make canceled a package based on pathvariable 'id'

	users := api.Group("/users")

	users.Get("/all", controllers.ListUsers)             // /api/users/all : Listing all users
	users.Post("/register", controllers.RegisterNewUser) // /api/users/register : Register new user via input JSON

	users.Post("/login", controllers.Login) // /api/users/login : Login user

	users.Get("/password-reset-code", controllers.SendPasswordResetCode) // /api/users/password-reset-code : Sending password reset code
	users.Post("/check-code", controllers.CheckResetCode)                // /api/users/check-code : Checking the code
	users.Post("/password-reset", controllers.ResetPassword)             // /api/users/password-reset : Resetting the password

	// From this point, all user endpoints are being authenticated
	users.Use(middleware.RequireJwtTokenAuth)

	users.Get("/:id", controllers.GetUserById) // /api/users/{id} : Get User by Id
	users.Put("/:id", controllers.EditUser)    // /api/users/{id} : Edit User

	users.Get("/get-accesslevel/:id", controllers.GetAccessLevel) // /api/users/get-accesslevel/{id} : Get the access level of the {id}. user
	users.Post("/set-accesslevel", controllers.SetAccessLevel)    // /api/users/set-accesslevel : Set the access level of the user {email, accesslevel}

	users.Get("/:id/packages", controllers.GetPackagesUnderUser) // /api/users/{id}/packages : Get all packages which the {id}. user sent
	users.Delete("/:id", controllers.DeleteUserById)             // api/users/{id} Delete user by id

	lockers := api.Group("/lockers")

	lockers.Get("/all", controllers.ListLockers) // api/lockers/all : Listing all lockers

	// From this point, all locker endpoints are being authenticated
	lockers.Use(middleware.RequireJwtTokenAuth)
	lockers.Post("/new", controllers.AddNewLocker)                  // /api/lockers/new : Add new locker via input json
	lockers.Get("/packages/:id", controllers.GetPackagesByLockerID) // /api/lockers/packages/{id} : Get all the information about the packages that are in the {id}. locker
	lockers.Get("/fullness/:id", controllers.GetFullness)           // /api/lockers/fullness/{id} : Get the fullness stats (cap, number of package, percentage) of the {id}. locker

	emissions := api.Group("/emissions")
	emissions.Get("", controllers.GetAllAndPerPackageEmission)
}
