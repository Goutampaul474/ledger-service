package services

import (
	"banking-ledger/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateAccount(name string, balance int64) (*models.Account, error) {
	id := uuid.New().String()
	createdAt := time.Now()

	_, err := s.PG.Exec(context.Background(),
		"INSERT INTO accounts (id, name, balance, created_at) VALUES ($1, $2, $3, $4)",
		id, name, balance, createdAt)
	if err != nil {
		return nil, err
	}

	return &models.Account{ID: id, Name: name, Balance: balance, CreatedAt: createdAt}, nil
}

func (s *Service) GetAccountByID(id string) (*models.Account, error) {
	row := s.PG.QueryRow(context.Background(),
		"SELECT id, name, balance, created_at FROM accounts WHERE id = $1", id)

	var acc models.Account
	if err := row.Scan(&acc.ID, &acc.Name, &acc.Balance, &acc.CreatedAt); err != nil {
		return nil, err
	}
	return &acc, nil
}
