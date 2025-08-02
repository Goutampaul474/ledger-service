package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() (*pgxpool.Pool, error) {
	dsn := "postgres://postgres:Goutam@123@localhost:5432/banking"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Postgres: %w", err)
	}

	// Run migration
	_, err = pool.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS accounts (
            id UUID PRIMARY KEY,
            name TEXT NOT NULL,
            balance BIGINT NOT NULL,
            created_at TIMESTAMP DEFAULT NOW()
        );
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Postgres: %w", err)
	}

	return pool, nil
}
