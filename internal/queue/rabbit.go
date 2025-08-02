package queue

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbit() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = ch.QueueDeclare(
		"transactions",
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
