package configs

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func ConnectToMongo() (*mongo.Client, error) {
	conn, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		return nil, err
	}
	client = conn

	return client, nil
}

func DisconnectFromMongo() error {
	if err := client.Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}

func GetMongoClient() *mongo.Client {
	return client
}
