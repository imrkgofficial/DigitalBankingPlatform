package database

import (
	"log"
	"os"
)

// LoadDatabaseConfig loads environment variables for MongoDB and MySQL.
func LoadDatabaseConfig() {
	mongoURI := os.Getenv("MONGO_URI")
	mysqlDSN := os.Getenv("MYSQL_DSN")

	// MongoDB and MySQL connection
	if mongoURI == "" || mysqlDSN == "" {
		log.Fatal("Database configuration not found in environment variables")
	}
}
