package main

import (
	"fmt"
	"log"

	"github.com/DDLD93/go-wallet-service/src/controllers"
	"github.com/DDLD93/go-wallet-service/src/connections"
	"github.com/DDLD93/go-wallet-service/src/models"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection string.
	dbConnectionString := "user=ujere password=123456 dbname=walletService sslmode=disable"

	// Create a new database connection.
	dbConnector, err := connections.NewDBConnector(dbConnectionString)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer dbConnector.DB.Close()

	// Create a DBConnector instance for the controllers.
	controllerDBConnector := &controllers.DBConnector{dbConnector}

	// Example usage:

	// Create a new wallet
	newWallet := &models.Wallet{ID: 1,UserID: 2, Balance: 100.0, Currency: "USD"}
	err = controllerDBConnector.CreateWallet(newWallet)
	if err != nil {
		log.Fatalf("Error creating wallet: %v", err)
	}

	// Create a debit transaction
	debitAmount := 25.0
	debitDescription := "Groceries"
	err = controllerDBConnector.CreateTransaction(newWallet, debitAmount, debitDescription, models.Debit)
	if err != nil {
		log.Fatalf("Error creating debit transaction: %v", err)
	}

	// Get the wallet's balance
	walletID := newWallet.ID
	balance, err := controllerDBConnector.GetWalletBalance(walletID)
	if err != nil {
		log.Fatalf("Error getting wallet balance: %v", err)
	}
	fmt.Printf("Wallet Balance: %.2f\n", balance)

	// Get transaction history for the wallet
	transactions, err := controllerDBConnector.GetWalletTransactionHistory(walletID)
	if err != nil {
		log.Fatalf("Error getting transaction history: %v", err)
	}
	fmt.Println("Transaction History:")
	for _, transaction := range transactions {
		fmt.Printf("ID: %d, Amount: %.2f, Description: %s\n", transaction.ID, transaction.Amount, transaction.Description)
	}

	// Get transaction history for a user
	userID := "your-user-id"
	userTransactions, err := controllerDBConnector.GetUserTransactionHistory(userID)
	if err != nil {
		log.Fatalf("Error getting user transaction history: %v", err)
	}
	fmt.Println("User Transaction History:")
	for _, transaction := range userTransactions {
		fmt.Printf("ID: %d, Amount: %.2f, Description: %s\n", transaction.ID, transaction.Amount, transaction.Description)
	}
}
