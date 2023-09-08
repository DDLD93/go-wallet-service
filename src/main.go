package main

import (
	"log"
	"time"

	"github.com/DDLD93/go-wallet-service/src/connections"
)

type WalletServiceServer struct {
	dbConnector *connections.DBConnector
}

func main() {
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

	// Continue with your application logic or start your gRPC server here
}
