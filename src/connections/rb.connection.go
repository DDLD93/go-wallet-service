package connections

import (
    "github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

func InitRabbitMQConnection(amqpURL string) error {
    conn, err := amqp.Dial(amqpURL)
    if err != nil {
        return err
    }

    RabbitMQConn = conn
    return nil
}
