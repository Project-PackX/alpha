package initializers

import (
	"PackX/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// A programban szereplő adatbázis definiálása
var DB *gorm.DB

// Adatbázishot csatlakozás
func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Nem sikerült kapcsolódni az adatbázishoz")
	}
}

// Automigrálás
func SyncDB() {
	DB.AutoMigrate(&models.Package{})
}

// ---------- TESZT JELLEGGEL -------------------
type TestPackage struct {
	ID        string  `json:"id"`
	Sender    string  `json:"sender"`
	Price     float32 `json:"price"`
	Delivered bool    `json:"delivered"`
}

var TestDB = []TestPackage{}
