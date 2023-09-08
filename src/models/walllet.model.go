package models

import "time"

// Wallet represents a user's wallet.
type Wallet struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

