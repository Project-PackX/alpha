package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Reading the .env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Nem sikerült beolvasni a .env fájlt")
	}
}
