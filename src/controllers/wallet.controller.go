package controllers

import (
	"database/sql"
	"fmt"

	"github.com/DDLD93/go-wallet-service/src/connections"
	"github.com/DDLD93/go-wallet-service/src/models"
	_ "github.com/lib/pq"
)
type DBConnector struct {
	*connections.DBConnector
}

// CreateWallet creates a new wallet in the database.
func (connector *DBConnector) CreateWallet(wallet *models.Wallet) error {
	// Start a new database transaction.
	tx, err := connector.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// Insert the wallet record within the transaction.
	_, err = tx.Exec("INSERT INTO wallets (user_id, balance, currency) VALUES ($1, $2, $3)",
		wallet.UserID, wallet.Balance, wallet.Currency)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit the transaction to make it atomic.
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetWalletBalance retrieves the current balance of a wallet from the database.
func (connector *DBConnector) GetWalletBalance(walletID int) (float64, error) {
	var balance float64

	// Execute a SQL query to retrieve the balance of the specified wallet.
	err := connector.DB.QueryRow("SELECT balance FROM wallets WHERE id = $1", walletID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("wallet not found") // Return a specific error message for no rows found.
		}
		return 0, err
	}

	return balance, nil
}


// CreateTransaction creates a debit transaction in the database and updates the wallet's balance.
func (connector *DBConnector) CreateTransaction(wallet *models.Wallet, amount float64, description string, transactionType models.TransactionType) error {
	if transactionType != models.Debit {
		return fmt.Errorf("invalid transaction type")
	}

	// Retrieve the current balance of the wallet from the database.
	currentBalance, err := connector.GetWalletBalance(wallet.ID)
	if err != nil {
		return err
	}

	// Check if the wallet balance is sufficient for the debit transaction.
	if currentBalance < amount {
		return fmt.Errorf("insufficient balance")
	}

	tx, err := connector.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// Insert the transaction record within the transaction.
	_, err = tx.Exec("INSERT INTO transactions (wallet_id, amount, description, type) VALUES ($1, $2, $3, $4)",
		wallet.ID, amount, description, transactionType)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Update the wallet's balance after the debit transaction.
	wallet.Balance -= amount

	// Update the wallet's balance in the database.
	_, err = tx.Exec("UPDATE wallets SET balance = $1 WHERE id = $2", wallet.Balance, wallet.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetTransactionHistory retrieves transaction history for a wallet.
func (connector *DBConnector) GetWalletTransactionHistory(walletID int) ([]models.Transaction, error) {
	rows, err := connector.DB.Query("SELECT id, wallet_id, amount, description, timestamp, type FROM transactions WHERE wallet_id = $1", walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.WalletID, &transaction.Amount, &transaction.Description, &transaction.Timestamp, &transaction.Type)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (connector *DBConnector) GetUserTransactionHistory(UserID string) ([]models.Transaction, error) {
	rows, err := connector.DB.Query("SELECT id, wallet_id, amount, description, timestamp, type FROM transactions WHERE user_id = $1", UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.WalletID, &transaction.Amount, &transaction.Description, &transaction.Timestamp, &transaction.Type)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
