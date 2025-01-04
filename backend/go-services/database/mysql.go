package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var MySQLDB *sql.DB

// ConnectMySQL establishes a connection to MySQL and returns the DB object.
func ConnectMySQL(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("MySQL connection failed:", err)
	}
	fmt.Println("Successfully connected to MySQL!")
	return db
}
