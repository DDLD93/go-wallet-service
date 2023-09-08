package models

import "time"

// TransactionType represents the type of transaction (debit or credit).
type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

// Transaction represents a financial transaction.
type Transaction struct {
	ID          int             `json:"id"`
	WalletID    string          `json:"wallet_id"`
	UserID      string          `json:"user_id"`
	Email       string          `json:"email"`
	Amount      float64         `json:"amount"`
	Description string          `json:"description"`
	Timestamp   time.Time       `json:"timestamp"`
	Type        TransactionType `json:"type"`
}
