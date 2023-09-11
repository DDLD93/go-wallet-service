package configs

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config represents the configuration of your application.
type Config struct {
    DBConnectionString string
    // Add other configuration variables as needed
}

// LoadConfig loads the configuration from environment variables or a .env file.
func LoadConfig() *Config {
    // Load environment variables from a .env file (if available)
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables directly.")
    }

    // Retrieve PostgreSQL connection parameters
    pgHost := getEnv("DB_HOST")
    pgPort := getEnv("DB_PORT")
    pgUser := getEnv("DB_USER")
    pgPassword := getEnv("DB_PASSWORD")
    pgDBName := getEnv("DB_DBNAME")

    // Construct the PostgreSQL connection string
    pgConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        pgHost, pgPort, pgUser, pgPassword, pgDBName)

    config := &Config{
        DBConnectionString: pgConnectionString,
        // Add other environment variables and their default values here
    }

    return config
}

// getEnv retrieves an environment variable or returns a default value if it's not set.
func getEnv(key string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Fatalf("ENV Variable %s not found.", key)
    }
    return value
}
