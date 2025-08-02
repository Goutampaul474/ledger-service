package models

import "time"

// PostgreSQL: Accounts
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// MongoDB: Transactions
type Transaction struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Type      string    `json:"type"` // deposit | withdraw
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
