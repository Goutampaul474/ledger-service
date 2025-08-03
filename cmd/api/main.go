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

	// Queue (RabbitMQ)
	_, ch, err := queue.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
	}

	svc := &services.Service{PG: pg, MQ: ch, MongoDB: mongoCol}
	
	h := &handlers.Handler{S: svc}
// Start RabbitMQ consumer
queue.StartConsumer(svc)

	r := gin.Default()

	// Routes
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.POST("/transactions", h.NewTransaction)
	r.GET("/accounts/:id/transactions", h.GetTransactionHistory)
	r.GET("/health", h.HealthCheck)

	log.Println("API running on :8080")
	r.Run(":8080")
}
