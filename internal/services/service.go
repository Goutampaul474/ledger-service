package services

import (
	"encoding/json"

	"banking-ledger/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	PG      *pgxpool.Pool
	MQ      *amqp.Channel
	MongoDB *mongo.Collection
}



func (s *Service) PublishTransaction(tx models.Transaction) error {
	body, err := json.Marshal(tx)
	if err != nil {
		return err
	}
	return s.MQ.Publish(
		"", "transactions", false, false,
		amqp.Publishing{ContentType: "application/json", Body: body},
	)
}
