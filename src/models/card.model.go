package models

import "time"

type Card struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	AuthorizationCode string    `json:"authorization_code"`
	BIN               string    `json:"bin"`
	Last4             string    `json:"last4"`
	ExpMonth          string    `json:"exp_month"`
	ExpYear           string    `json:"exp_year"`
	Channel           string    `json:"channel"`
	CardType          string    `json:"card_type"`
	Bank              string    `json:"bank"`
	CountryCode       string    `json:"country_code"`
	Brand             string    `json:"brand"`
	Reusable          bool      `json:"reusable"`
	Signature         string    `json:"signature"`
	AccountName       string    `json:"account_name"`
	Timestamp         time.Time `json:"timestamp"`
}
