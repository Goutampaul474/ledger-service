package main

import (
	"banking-ledger/internal/db"
	"banking-ledger/internal/models"
	"banking-ledger/internal/queue"
	"banking-ledger/internal/services"
	"context"
	"encoding/json"
	"log"
)

func main() {
	// DB connections
	pg, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	mongoClient, mongoCol, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(nil)

	// RabbitMQ
	conn, ch, err := queue.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	svc := &services.Service{PG: pg, MQ: ch, MongoDB: mongoCol}

	msgs, err := ch.Consume(
		"transactions", // queue
		"",             // consumer name
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Worker started. Listening for transactions...")

	// process messages
	for d := range msgs {
		var txn models.Transaction
		if err := json.Unmarshal(d.Body, &txn); err != nil {
			log.Printf("Invalid transaction JSON: %v", err)
			continue
		}

		// process in Postgres
		err := svc.ProcessTransaction(&txn)
		if err != nil {
			log.Printf("Failed to process txn %s: %v", txn.ID, err)
			txn.Status = "failed"
		} else {
			txn.Status = "success"
		}

		// log in MongoDB
		_, err = svc.MongoDB.InsertOne(context.Background(), txn)
		if err != nil {
			log.Printf("Failed to insert txn into MongoDB: %v", err)
		}

		log.Printf("Processed transaction: %+v", txn)
	}
}
