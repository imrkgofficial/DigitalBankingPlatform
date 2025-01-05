package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/imrkgofficial/DigitalBankingPlatform/backend/go-services/database"
	"github.com/imrkgofficial/DigitalBankingPlatform/backend/go-services/authentication"
	"github.com/imrkgofficial/DigitalBankingPlatform/backend/go-services/account"
	"github.com/imrkgofficial/DigitalBankingPlatform/backend/go-services/transaction"
)

func init() {
	// Load environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Optional: Check if required environment variables are set
	requiredEnvVars := []string{"MONGO_URI", "MYSQL_DSN"}
	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}
}

// monitorConnections checks MongoDB and MySQL connections periodically
func monitorConnections(mongoURI string, mysqlDSN string) {
	for {
		// Check MongoDB connection
		mongoClient := database.ConnectMongoDB(mongoURI)
		err := mongoClient.Ping(nil, nil)
		if err != nil {
			log.Printf("[WARNING] [%s]: MongoDB connection failed: %v", time.Now().Format(time.RFC3339), err)
		} else {
			log.Printf("[INFO] [%s]: Successfully connected to MongoDB!", time.Now().Format(time.RFC3339))
		}
		mongoClient.Disconnect(nil)

		// Check MySQL connection
		mysqlDB := database.ConnectMySQL(mysqlDSN)
		err = mysqlDB.Ping()
		if err != nil {
			log.Printf("[WARNING] [%s]: MySQL connection failed: %v", time.Now().Format(time.RFC3339), err)
		} else {
			log.Printf("[INFO] [%s]: Successfully connected to MySQL!", time.Now().Format(time.RFC3339))
		}
		mysqlDB.Close()

		// Sleep for 10 seconds before rechecking
		time.Sleep(10 * time.Second)
	}
}

func main() {
	// Load database configurations from environment variables
	databaseURI := os.Getenv("MONGO_URI")
	mysqlDSN := os.Getenv("MYSQL_DSN")

	// Connect to MongoDB
	mongoClient := database.ConnectMongoDB(databaseURI)
	defer func() {
		if mongoClient != nil {
			mongoClient.Disconnect(nil)
		}
	}()

	// Check MongoDB connection
	if err := mongoClient.Ping(nil, nil); err != nil {
		log.Fatalf("FATAL [%s]: MongoDB connection failed: %v", time.Now().Format(time.RFC3339), err)
	} else {
		fmt.Println("Successfully connected to MongoDB!") // Use of fmt
	}

	// Connect to MySQL
	mysqlDB := database.ConnectMySQL(mysqlDSN)
	defer mysqlDB.Close()

	// Check MySQL connection
	if err := mysqlDB.Ping(); err != nil {
		log.Fatalf("FATAL [%s]: MySQL connection failed: %v", time.Now().Format(time.RFC3339), err)
	} else {
		fmt.Println("Successfully connected to MySQL!") // Use of fmt
	}

	// Launch a separate goroutine to monitor database connections
	go monitorConnections(databaseURI, mysqlDSN)

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
	fmt.Println("Server started at http://localhost:8080") // Use of fmt
	log.Fatal(server.ListenAndServe())
}
