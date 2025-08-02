package main

import (
	"banking-ledger/internal/db"
	"banking-ledger/internal/models"
	"banking-ledger/internal/queue"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	pg, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	mongoClient, mongoCol, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(nil)

	_, ch, err := queue.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume("transactions", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Worker listening for transactions...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var tx models.Transaction
			if err := json.Unmarshal(d.Body, &tx); err != nil {
				log.Println("Invalid message:", err)
				continue
			}

			// Apply transaction
			var newBalance int64
			if tx.Type == "deposit" {
				err = pg.QueryRow(context.Background(),
					"UPDATE accounts SET balance = balance + $1 WHERE id=$2 RETURNING balance",
					tx.Amount, tx.AccountID).Scan(&newBalance)
			} else if tx.Type == "withdraw" {
				err = pg.QueryRow(context.Background(),
					"UPDATE accounts SET balance = balance - $1 WHERE id=$2 AND balance >= $1 RETURNING balance",
					tx.Amount, tx.AccountID).Scan(&newBalance)
			}

			if err != nil {
				tx.Status = "failed"
			} else {
				tx.Status = "success"
			}

			_, err = mongoCol.InsertOne(context.Background(), bson.M{
				"transaction_id": tx.ID,
				"account_id":     tx.AccountID,
				"type":           tx.Type,
				"amount":         tx.Amount,
				"status":         tx.Status,
				"created_at":     tx.CreatedAt,
			})
			if err != nil {
				log.Println("Mongo insert failed:", err)
			}
		}
	}()

	<-forever
}
