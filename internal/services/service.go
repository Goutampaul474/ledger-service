package services

import (
	"context"
	"encoding/json"
	"time"

	"banking-ledger/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	PG      *pgxpool.Pool
	MQ      *amqp.Channel
	MongoDB *mongo.Collection
}

func (s *Service) CreateAccount(ctx context.Context, name string, balance int64) (*models.Account, error) {
	acc := &models.Account{
		ID:        uuid.NewString(),
		Name:      name,
		Balance:   balance,
		CreatedAt: time.Now(),
	}

	_, err := s.PG.Exec(ctx, `INSERT INTO accounts (id, name, balance, created_at) VALUES ($1, $2, $3, $4)`,
		acc.ID, acc.Name, acc.Balance, acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
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
