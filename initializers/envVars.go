package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Reading the .env file
func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
}
