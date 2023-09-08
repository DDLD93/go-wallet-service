package configs

import (
	"fmt"
	"os"
)

// Config represents the configuration of your application.
type Config struct {
    DBConnectionString string
    // Add other configuration variables as needed
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
    config := &Config{
        DBConnectionString: getEnv("DB_CONNECTION_STRING", "user=ujere password=123456 dbname=postgres sslmode=disable"),
        // Add other environment variables and their default values here
    }

    return config
}

// getEnv retrieves an environment variable or returns a default value if it's not set.
func getEnv(key, defaultValue string) string {
	
    value, exists := os.LookupEnv(key)
	fmt.Print("env >>>>>>")

	fmt.Print(value)
    if !exists {
        return defaultValue
    }
    return value
}
