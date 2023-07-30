package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// A .env fájl beolvasása
func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
}
