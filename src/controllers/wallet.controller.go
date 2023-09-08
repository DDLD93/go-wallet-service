package controllers

import (
    "database/sql"

    "github.com/DDLD93/go-wallet-service/src/connections"
    "github.com/DDLD93/go-wallet-service/src/models"
    _ "github.com/lib/pq"
)

type DBConnector struct {
    *connections.DBConnector
}

type Response struct {
    Ok      bool
    Data    interface{}
    Message string
}

// CreateWallet creates a new wallet in the database.
func (connector *DBConnector) CreateWallet(wallet *models.Wallet) Response {
    response := Response{}

    // Start a new database transaction.
    tx, err := connector.DB.Begin()
    if err != nil {
        response.Message = err.Error()
        return response
    }
    defer func() {
        if p := recover(); p != nil {
            _ = tx.Rollback()
            panic(p)
        }
    }()

    // Insert the wallet record within the transaction.
    _, err = tx.Exec("INSERT INTO wallets (user_id,email, balance, currency) VALUES ($1,$2, $3, $4)",
        wallet.UserID, wallet.Email, wallet.Balance, wallet.Currency)
    if err != nil {
        _ = tx.Rollback()
        response.Message = err.Error()
        return response
    }

    // Commit the transaction to make it atomic.
    err = tx.Commit()
    if err != nil {
        response.Message = err.Error()
        return response
    }

    response.Ok = true
    return response
}

// GetWalletBalance retrieves the current balance of a wallet from the database.
func (connector *DBConnector) GetWalletBalance(UserID string) Response {
    response := Response{}

    var balance float64

    // Execute a SQL query to retrieve the balance of the specified wallet.
    err := connector.DB.QueryRow("SELECT balance FROM wallets WHERE user_id = $1", UserID).Scan(&balance)
    if err != nil {
        if err == sql.ErrNoRows {
            response.Message = "Wallet not found"
        } else {
            response.Message = err.Error()
        }
        return response
    }

    response.Ok = true
    response.Data = balance
    return response
}

// CreateTransaction creates a debit transaction in the database and updates the wallet's balance.
func (connector *DBConnector) CreateTransaction(userID string, email string, amount float64, description string, transactionType models.TransactionType) Response {
    response := Response{}

    if transactionType != models.Debit {
        response.Message = "Invalid transaction type"
        return response
    }

    // Retrieve the current balance of the wallet from the database.
    currentBalanceResponse := connector.GetWalletBalance(userID)
    if !currentBalanceResponse.Ok {
        response.Message = currentBalanceResponse.Message
        return response
    }

    currentBalance := currentBalanceResponse.Data.(float64)

    // Check if the wallet balance is sufficient for the debit transaction.
    if currentBalance < amount {
        response.Message = "Insufficient balance"
        return response
    }

    tx, err := connector.DB.Begin()
    if err != nil {
        response.Message = err.Error()
        return response
    }
    defer func() {
        if p := recover(); p != nil {
            _ = tx.Rollback()
            panic(p)
        }
    }()

    // Insert the transaction record within the transaction.
    _, err = tx.Exec("INSERT INTO transactions (user_id,email ,amount, description, type) VALUES ($1, $2, $3, $4,$5)",
        userID, email, amount, description, transactionType)
    if err != nil {
        _ = tx.Rollback()
        response.Message = err.Error()
        return response
    }

    // Update the wallet's balance after the debit transaction.
    newBalance := currentBalance - amount

    // Update the wallet's balance in the database.
    _, err = tx.Exec("UPDATE wallets SET balance = $1 WHERE user_id = $2", newBalance, userID)
    if err != nil {
        _ = tx.Rollback()
        response.Message = err.Error()
        return response
    }

    err = tx.Commit()
    if err != nil {
        response.Message = err.Error()
        return response
    }

    response.Ok = true
    return response
}

// GetWalletTransactionHistory retrieves transaction history for a wallet.
func (connector *DBConnector) GetWalletTransactionHistory(userID string) Response {
    response := Response{}

    rows, err := connector.DB.Query("SELECT id, amount, description, timestamp, type FROM transactions WHERE user_id = $1", userID)
    if err != nil {
        response.Message = err.Error()
        return response
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var transaction models.Transaction
        err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Description, &transaction.Timestamp, &transaction.Type)
        if err != nil {
            response.Message = err.Error()
            return response
        }
        transactions = append(transactions, transaction)
    }

    if err := rows.Err(); err != nil {
        response.Message = err.Error()
        return response
    }

    response.Ok = true
    response.Data = transactions
    return response
}
