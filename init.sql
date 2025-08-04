CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
