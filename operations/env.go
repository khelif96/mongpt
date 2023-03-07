package operations

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariablesIntoEnvironment() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
