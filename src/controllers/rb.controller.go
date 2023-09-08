package controllers

import (
	"log"

	"github.com/DDLD93/go-wallet-service/src/connections"
	"github.com/streadway/amqp"
	// Update with your actual app package name
)

func PublishMessage(exchangeName, routingKey string, message []byte) error {
    ch, err := connections.RabbitMQConn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    err = ch.Publish(
        exchangeName,
        routingKey,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        message,
        },
    )
    if err != nil {
        return err
    }

    log.Printf("Message sent: %s", string(message))
    return nil
}

func ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
    ch, err := connections.RabbitMQConn.Channel()
    if err != nil {
        return nil, err
    }

    msgs, err := ch.Consume(
        queueName,
        "",
        true,  // Auto-acknowledge messages
        false, // Exclusive
        false, // No-local
        false, // No-wait
        nil,   // Args
    )
    if err != nil {
        return nil, err
    }

    log.Printf("Started consuming messages from queue: %s", queueName)
    return msgs, nil
}