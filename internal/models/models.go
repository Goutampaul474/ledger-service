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
    ID        string    `json:"id" bson:"id"`
    AccountID string    `json:"account_id" bson:"accountid"`
    Type      string    `json:"type" bson:"type"`
    Amount    int64     `json:"amount" bson:"amount"`
    Status    string    `json:"status" bson:"status"`
    CreatedAt time.Time `json:"created_at" bson:"createdat"`
}

