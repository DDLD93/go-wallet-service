package models

import "time"

// Wallet represents a user's wallet.
type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

// type TransactionType string

// const (
//     Debit  TransactionType = "debit"
//     Credit TransactionType = "credit"
// )

// // Transaction represents a financial transaction.
// type Transaction struct {
//     ID          int            `json:"id"`
//     WalletID    int            `json:"wallet_id"`
//     Amount      float64        `json:"amount"`
//     Description string         `json:"description"`
//     Timestamp   time.Time      `json:"timestamp"`
//     Type        TransactionType `json:"type"`
// }
