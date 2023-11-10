package main

import (
	"os"

	"github.com/Project-PackX/backend/initializers"
	"github.com/Project-PackX/backend/utils"
	"github.com/mikhail-bigun/fiberlogrus"

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
	// Creating the fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	logger := utils.Logger
	logger.Info("App successfully started on port " + os.Getenv("PORT"))

	app.Use(
		fiberlogrus.New(
			fiberlogrus.Config{
				Logger: logger,
				Tags:   fiberlogrus.CommonTags,
			}))

	// Setting up the Cross-Origin Resource Sharing config
	app.Use(cors.New(cors.ConfigDefault))

	// Endpoints management via function
	Routes(app)

	// Start the webserver
	app.Listen(":" + os.Getenv("PORT"))
}
