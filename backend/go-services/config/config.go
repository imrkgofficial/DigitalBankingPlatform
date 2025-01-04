package config

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadConfig loads configuration variables from the environment.
func LoadConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
