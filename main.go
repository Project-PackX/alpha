package main

import (
	"PackX/initializers"
	"PackX/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()  // Loading environment variables (port, database)
	initializers.ConnectToDatabase() // Conencting to database based on env vars

	// FOR TESTING PURPOSES
	initializers.DropTables()
	// ------

	initializers.SyncDB() // Creating tanles based on the models

	// FOR TESTING PURPOSES
	initializers.GenerateTestEntries() // Generating test datas
	// ------

}

func main() {
	// Set up views
	engine := html.New("./views", ".html")

	// Creating the fiber app with views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure the application
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	// Endpoints management via function
	Routes(app)

	// Start the webserver
	app.Listen(":" + os.Getenv("PORT"))

}
