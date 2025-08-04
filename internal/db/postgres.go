package db

import (
	"banking-ledger/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() (*pgxpool.Pool, error) {
	 dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        config.PostgresUser,
        config.PostgresPassword,
        config.PostgresHost,
        config.PostgresPort,
        config.PostgresDB,
    )
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
	log.Println("Connected to Postgres")

	return pool, nil
}
