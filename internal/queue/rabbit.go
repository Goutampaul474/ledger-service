package queue

import (
	"banking-ledger/internal/config"
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbit() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(config.RabbitURI)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = ch.QueueDeclare(
		config.TransactionsQ,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	log.Println("Connected to RabbitMQ")
	return conn, ch, nil
}
