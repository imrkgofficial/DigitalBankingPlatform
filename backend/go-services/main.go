package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"DigitalBankingPlatform/backend/go-services/database"
	"DigitalBankingPlatform/backend/go-services/authentication"
	"DigitalBankingPlatform/backend/go-services/account"
	"DigitalBankingPlatform/backend/go-services/transaction"
)

func init() {
	// Load environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Load database configurations
	databaseURI := os.Getenv("MONGO_URI")
	mysqlDSN := os.Getenv("MYSQL_DSN")

	// Connect to MongoDB and MySQL
	mongoClient := database.ConnectMongoDB(databaseURI)
	defer mongoClient.Disconnect(nil)
	mysqlDB := database.ConnectMySQL(mysqlDSN)
	defer mysqlDB.Close()

	// Log successful connections
	fmt.Println("Successfully connected to MongoDB and MySQL!")

	// Create a new router
	r := mux.NewRouter()

	// Define routes for authentication, accounts, and transactions
	r.HandleFunc("/api/auth/login", authentication.LoginHandler).Methods("POST")
	r.HandleFunc("/api/auth/register", authentication.RegisterHandler).Methods("POST")

	r.HandleFunc("/api/accounts", account.GetAccountsHandler).Methods("GET")
	r.HandleFunc("/api/accounts/{id}", account.GetAccountHandler).Methods("GET")
	r.HandleFunc("/api/accounts/{id}/transfer", account.TransferFundsHandler).Methods("POST")

	r.HandleFunc("/api/transactions", transaction.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/api/transactions/{id}", transaction.GetTransactionHandler).Methods("GET")

	// Set up server with 10-second timeout to avoid hanging
	server := &http.Server{
		Addr:           ":8080", // Port on which server will run
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
