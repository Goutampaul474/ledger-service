package services

import (
	"banking-ledger/internal/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) GetTransactionsByAccountID(accountID string) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := s.MongoDB.Find(ctx, bson.M{"accountid": accountID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var txns []models.Transaction
	if err := cur.All(ctx, &txns); err != nil {
		return nil, err
	}
	return txns, nil
}

func (s *Service) ProcessTransaction(txn *models.Transaction) error {
	ctx := context.Background()

	switch txn.Type {
	case "deposit":
		_, err := s.PG.Exec(ctx,
			"UPDATE accounts SET balance = balance + $1 WHERE id = $2",
			txn.Amount, txn.AccountID)
		return err

	case "withdraw":
		// Only allow if balance >= amount
		_, err := s.PG.Exec(ctx,
			"UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1",
			txn.Amount, txn.AccountID)
		return err

	default:
		return fmt.Errorf("invalid transaction type: %s", txn.Type)
	}
}

