package connections

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DBConnector represents the database connection.
type DBConnector struct {
	DB *sql.DB
}

// NewDBConnector creates a new database connection and initializes the wallet table.
func NewDBConnector(connectionString string) (*DBConnector, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Initialize the wallet table if it doesn't exist.
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return &DBConnector{DB: db}, nil
}



// createTables creates necessary tables if they don't exist.
func createTables(db *sql.DB) error {
	// Create the wallets table if it doesn't exist.
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS wallets (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			balance NUMERIC(10, 2) NOT NULL,
			currency VARCHAR(3) NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating wallets table: %v", err)
		return err
	}

	// Create the transactions table if it doesn't exist.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			wallet_id INT NOT NULL,
			amount NUMERIC(10, 2) NOT NULL,
			description VARCHAR(255),
			timestamp TIMESTAMPTZ DEFAULT current_timestamp,
			type VARCHAR(10) NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating transactions table: %v", err)
		return err
	}

	// Create the cards table if it doesn't exist.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cards (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			card_number VARCHAR(16) NOT NULL,
			expiration_date DATE NOT NULL,
			cvv VARCHAR(3) NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating cards table: %v", err)
		return err
	}

	// Add more table creation statements if needed.

	return nil
}

