package databases

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectToMongoDB() error {
	url := os.Getenv("MONGODB_URL")

	if url == "" {
		return errors.New("the 'MONGODB_URL' environmental variable must be set")
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.Background(), nil)

	if err != nil {
		return errors.New("can't verify a connection to MongoDB")
	}

	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	database := os.Getenv("DATABASE")

	return mongoClient.Database(database).Collection(collectionName)
}

func DisconnectFromMongoDB() {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
