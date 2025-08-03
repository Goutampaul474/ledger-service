# 🏦 Banking Ledger Service (Golang)

A backend service written in **Go** that simulates a simple banking ledger.  
It manages **accounts** (in Postgres), **transactions** (in MongoDB), and uses **RabbitMQ** for asynchronous processing.

---

## 🚀 Features

- Create bank accounts with an initial balance (**Postgres**)
- Deposit / Withdraw money using transaction requests (**RabbitMQ** + Worker)
- Maintain an immutable transaction log (**MongoDB**)
- Ensure consistency of balances (ACID‑like)
- REST APIs built with **Gin**
- Asynchronous transaction processing for scalability
- Health check endpoint for service monitoring

---

## 🛠️ Tech Stack

- **Golang** (Gin framework)
- **Postgres** → Stores accounts and balances
- **MongoDB** → Stores transaction logs
- **RabbitMQ** → Queue for transaction requests
- **pgxpool** → PostgreSQL driver
- **mongo-driver** → MongoDB driver
- **amqp** → RabbitMQ client

---

## 📂 Project Structure

banking-ledger/
│── cmd/
│ ├── api/ # API server entrypoint
│ │ └── main.go
│ └── worker/ # Worker to process transactions
│ └── main.go
│
│── internal/
│ ├── db/ # Database connections
│ │ ├── postgres.go
│ │ └── mongo.go
│ │
│ ├── handlers/ # Gin route handlers
│ │ ├── account.go
│ │ ├── transaction.go
│ │ └── health.go
│ │
│ ├── models/ # Data models
│ │ └── models.go
│ │
│ ├── services/ # Business logic
│ │ ├── account.go
│ │ ├── transaction.go
│ │ └── service.go
│ │
│ └── queue/ # RabbitMQ consumer
|   ├── rabbitmq.go
│   └── consumer.go
│
└── go.mod


## ⚙️ Setup Instructions

### 1. Install Dependencies
- [PostgreSQL](https://www.postgresql.org/download/) (default port: `5432`)
- [MongoDB](https://www.mongodb.com/try/download/community) (default port: `27017`)
- [RabbitMQ](https://www.rabbitmq.com/download.html) (default port: `5672`, UI: `15672`)

### 2. Create Postgres Database
```sql
CREATE DATABASE banking;

\c banking

CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
3. Start MongoDB
Use MongoDB Compass or mongod CLI

Database: banking

Collection: transactions (auto-created when first transaction is inserted)

4. Start RabbitMQ

# Windows Service
rabbitmq-service start

# Enable Management UI
rabbitmq-plugins enable rabbitmq_management
Visit → http://localhost:15672 (user: guest, pass: guest)

▶️ Running the Application
Start API Server

go run ./cmd/api
Expected log:


Connected to Postgres
Connected to MongoDB
Connected to RabbitMQ
API running on :8080
Start Worker

go run ./cmd/worker
Expected log:



Connected to Postgres
Connected to MongoDB
Connected to RabbitMQ
Worker started. Listening for transactions...
🌐 API Endpoints
1. Create Account
POST /accounts


{
  "name": "Alice",
  "balance": 1000
}
Response:


{
  "id": "5ff465c6-50e1-4a35-a349-386e4dd853d5",
  "name": "Alice",
  "balance": 1000,
  "created_at": "2025-08-03T21:08:42.1672407+05:30"
}
2. Get Account by ID
GET /accounts/:id

Response:


{
  "id": "5ff465c6-50e1-4a35-a349-386e4dd853d5",
  "name": "Alice",
  "balance": 1150,
  "created_at": "2025-08-03T21:08:42.1672407+05:30"
}
3. Submit Transaction (Deposit / Withdraw)
POST /transactions


{
  "account_id": "5ff465c6-50e1-4a35-a349-386e4dd853d5",
  "type": "deposit",
  "amount": 200
}
Response:


{
  "message": "Transaction submitted",
  "transaction_id": "abc123"
}
4. Get Transaction History
GET /accounts/:id/transactions

Response:


[
  {
    "id": "fad7e3b1-1faf-4fed-b706-e02efee66d17",
    "account_id": "5ff465c6-50e1-4a35-a349-386e4dd853d5",
    "type": "deposit",
    "amount": 200,
    "status": "success",
    "created_at": "2025-08-03T16:31:28.309Z"
  },
  {
    "id": "e4b9b3f3-9833-4ddd-bfcd-6c7a82f099cc",
    "account_id": "5ff465c6-50e1-4a35-a349-386e4dd853d5",
    "type": "withdraw",
    "amount": 50,
    "status": "success",
    "created_at": "2025-08-03T16:31:50.034Z"
  }
]
5. Health Check
GET /health

Response:


{
  "status": "ok",
  "postgres": true,
  "mongodb": true,
  "rabbitmq": true
}
🧪 Testing with curl

# Create account
curl -X POST http://localhost:8080/accounts \
     -H "Content-Type: application/json" \
     -d '{"name":"Alice","balance":1000}'

# Deposit
curl -X POST http://localhost:8080/transactions \
     -H "Content-Type: application/json" \
     -d '{"account_id":"<account_id>","type":"deposit","amount":200}'

# Withdraw
curl -X POST http://localhost:8080/transactions \
     -H "Content-Type: application/json" \
     -d '{"account_id":"<account_id>","type":"withdraw","amount":50}'

# Get account details
curl http://localhost:8080/accounts/<account_id>

# Get transaction history
curl http://localhost:8080/accounts/<account_id>/transactions

# Health check
curl http://localhost:8080/health
📖 Notes
Postgres ensures current balances are always consistent.

MongoDB acts as an immutable ledger for all transactions.

RabbitMQ decouples transaction requests from processing → improves scalability.

Run API + Worker separately for full functionality.


✅ Next Steps
Add unit tests (mocking Postgres, MongoDB, RabbitMQ).

Add Docker + docker-compose for one‑command setup.

Add retry logic for failed transactions.