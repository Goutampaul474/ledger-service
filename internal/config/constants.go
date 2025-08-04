package config

const (
    // Postgres
    PostgresUser     = "postgres"
    PostgresPassword = "password"
    PostgresDB       = "banking"
    PostgresHost     = "postgres"
    PostgresPort     = "5432"

    // PostgresUser     = "postgres"
    // PostgresPassword = "Goutam@123"
    // PostgresDB       = "banking"
    // PostgresHost     = "localhost"
    // PostgresPort     = "5432"

    // MongoDB
    MongoURI  = "mongodb://mongo:27017"
//    MongoURI  = "mongodb://localhost:27017"

    MongoDB   = "banking"
    MongoColl = "transactions"

    // RabbitMQ
    RabbitURI      = "amqp://guest:guest@rabbitmq:5672/"
    // RabbitURI="amqp://guest:guest@localhost:5672/"
    TransactionsQ  = "transactions"
)
