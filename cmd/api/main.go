package main

import (
	"banking-ledger/internal/db"
	"banking-ledger/internal/handlers"
	"banking-ledger/internal/queue"
	"banking-ledger/internal/services"
	"log"

	"github.com/gin-gonic/gin"
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

	// Queue
	_, ch, err := queue.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
	}

	svc := &services.Service{PG: pg, MQ: ch, MongoDB: mongoCol}
	h := &handlers.Handler{S: svc}

	r := gin.Default()
	r.POST("/accounts", h.CreateAccount)
	r.POST("/transactions", h.NewTransaction)

	log.Println("API running on :8080")
	r.Run(":8080")
}
