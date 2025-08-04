package queue

import (
	"banking-ledger/internal/models"
	"banking-ledger/internal/services"
	"context"
	"encoding/json"
	"log"
)

func StartConsumer(s *services.Service) {
	msgs, err := s.MQ.Consume(
		"transactions", // queue name
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var txn models.Transaction
			if err := json.Unmarshal(d.Body, &txn); err != nil {
				log.Printf("Invalid transaction: %v", err)
				continue
			}

			// Process transaction: update Postgres balance
			err := s.ProcessTransaction(&txn)
			if err != nil {
				log.Printf("Failed to process txn %s: %v", txn.ID, err)
				txn.Status = "failed"
			} else {
				txn.Status = "success"
			}

			// Save to MongoDB ledger
			_, err = s.MongoDB.InsertOne(context.Background(), txn)
			if err != nil {
				log.Printf("Failed to log txn in MongoDB: %v", err)
			}

			log.Printf("Processed transaction %+v", txn)
		}
	}()
}
