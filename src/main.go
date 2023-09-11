package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/DDLD93/go-wallet-service/src/configs"
	"github.com/DDLD93/go-wallet-service/src/connections"
	"github.com/DDLD93/go-wallet-service/src/controllers/paystack"
	"github.com/gin-gonic/gin"

	// grpc "github.com/DDLD93/go-wallet-service/src/gprc"
	_ "github.com/lib/pq"
)

type WalletServiceServer struct {
    dbConnector *connections.DBConnector
}
const (
    paystackSecretKey = "YOUR_PAYSTACK_SECRET_KEY"
    callbackPath      = "/paystack/callback" // Define your callback path
)

func main() {


    var pgDB *sql.DB
    var err error
     configs := configs.LoadConfig()

    // Continuously attempt to connect to PostgreSQL
    for {
        // Open the PostgreSQL database connection
        pgDB, err = sql.Open("postgres", configs.DBConnectionString)
        if err != nil {
            log.Printf("Failed to connect to PostgreSQL: %v", err)
            // Sleep for a while before attempting to reconnect
            time.Sleep(5 * time.Second)
            continue // Continue the loop to retry
        }

        // Test the PostgreSQL connection
        err = pgDB.Ping()
        if err != nil {
            log.Printf("Failed to ping PostgreSQL: %v", err)
            pgDB.Close() // Close the failed connection
            // Sleep for a while before attempting to reconnect
            time.Sleep(5 * time.Second)
            continue // Continue the loop to retry
        }

        log.Println("Connected to PostgreSQL!")
        break // Break the loop if the connection is successful
    }

    // RabbitMQ connection parameters
    amqpURL := "amqp://wallet:wallet@localhost:5672/"
    for {
        // Attempt to initialize RabbitMQ connection
        err := connections.InitRabbitMQConnection(amqpURL)
        if err != nil {
            log.Printf("Failed to initialize RabbitMQ connection: %v", err)
            // Sleep for a while before attempting to reconnect
            time.Sleep(5 * time.Second)
            continue // Continue the loop to retry
        }

        log.Println("RabbitMQ connection successfully initialized")
        // If the connection is successful, you can break the loop and proceed
        break
    }
        // Create a Gin router
        router := gin.Default()

        // Define a route for the Paystack callback
        router.POST(callbackPath, paystack.HandlePaystackCallback)
        log.Printf("Server is running on :8080%s", callbackPath)
        if err := router.Run(":8088"); err != nil {
            log.Fatal(err)
        }
    
        // Start the HTTP server
        // go func() {
        //     log.Printf("Server is running on :8080%s", callbackPath)
        //     if err := router.Run(":8088"); err != nil {
        //         log.Fatal(err)
        //     }
        // }()

}
