package db

import (
	"banking-ledger/internal/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(config.MongoDB).Collection(config.MongoColl)
	log.Println("Connected to MongoDB")
	return client, collection, nil
}
