package services

import "banking-ledger/internal/models"

// Defines what the Handler needs, not how it's implemented
type ServiceInterface interface {
    CreateAccount(name string, balance int64) (*models.Account, error)
    GetAccountByID(id string) (*models.Account, error)
    PublishTransaction(tx models.Transaction) error
    GetTransactionsByAccountID(id string) ([]models.Transaction, error)

    // Add these just for health checks
    IsPostgresHealthy() bool
    IsMongoHealthy() bool
    IsRabbitMQHealthy() bool
}

func (s *Service) IsPostgresHealthy() bool {
    return s.PG != nil
}
func (s *Service) IsMongoHealthy() bool {
    return s.MongoDB != nil
}
func (s *Service) IsRabbitMQHealthy() bool {
    return s.MQ != nil
}
