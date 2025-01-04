package transaction

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// GetTransactionsHandler returns all transactions.
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to retrieve all transactions
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of all transactions")
}

// GetTransactionHandler returns a single transaction by ID.
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID := params["id"]
	// Logic to retrieve transaction by ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Details for transaction ID: %s", transactionID)
}
