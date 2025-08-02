package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("banking").Collection("transactions")
	return client, collection, nil
}
