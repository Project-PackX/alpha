package main

import (
	"PackX/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()  // Loading environment variables (port, database)
	initializers.ConnectToDatabase() // Connecting to database based on env vars

	// FOR TESTING PURPOSES
	initializers.DropTables()
	// ------

	initializers.SyncDB() // Creating tanles based on the models

	// FOR TESTING PURPOSES
	initializers.GenerateTestEntries() // Generating test datas
	// ------

}

func main() {
	// Creating the fiber app with views
	app := fiber.New(fiber.Config{})

	// Setting up the Cross-Origin Resource Sharing config
	app.Use(cors.New(cors.ConfigDefault))
	/*
		Inside ConfigDefault:
			Next:         nil,
		    AllowOriginsFunc: nil,
		    AllowOrigins: "*",
		    AllowMethods: strings.Join([]string{
		       	fiber.MethodGet,
		       	fiber.MethodPost,
		       	fiber.MethodHead,
		       	fiber.MethodPut,
		       	fiber.MethodDelete,
		       	fiber.MethodPatch,
		    }, ","),
		    AllowHeaders:     "",
		    AllowCredentials: false,
			ExposeHeaders:    "",
		    MaxAge:           0,
	*/

	// Configure the application
	app.Static("/", "./public")

	// Endpoints management via function
	Routes(app)

	// Start the webserver
	app.Listen(":" + os.Getenv("PORT"))

}
