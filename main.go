package main

import (
	"PackX/initializers"
	"PackX/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()  // Környezeti változók betöltése (port, adatbázis)
	initializers.ConnectToDatabase() // Környezeti változók alapján csatlakozás az adatbázishoz
	initializers.SyncDB()            // Adatbázis adatok automigrálása a gorm DB-be
}

func main() {
	// Nézetek beállítása
	engine := html.New("./views", ".html")

	// Fiber inicializálása + nézet beállítás
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// 'app' beállítása
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	// Kölünböző végpontok és hozzájuk tartozó kezelőfüggvények kihelyezése
	Routes(app)

	// Webszerver elindítása
	app.Listen(":" + os.Getenv("PORT"))
}
