package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// A .env fájl beolvasása
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Nem sikerült beolvasni a .env fájlt")
	}
}
