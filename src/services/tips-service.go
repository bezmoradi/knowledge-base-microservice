package services

import (
	"context"
	"log"

	"github.com/bezmoradi/knowledge-base-microservice/src/databases"
	"github.com/bezmoradi/knowledge-base-microservice/src/helpers"
	"github.com/bezmoradi/knowledge-base-microservice/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TIPS_COLLECTION = "tips"

func createTip(path string) {
	collection := databases.GetCollection(TIPS_COLLECTION)
	doc := models.TipDocument{
		ID:   primitive.NewObjectID(),
		Path: path}

	collection.InsertOne(context.Background(), doc)
}

func deleteTip(id interface{}) {
	collection := databases.GetCollection(TIPS_COLLECTION)
	collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
}

func getRandomContents(count int) []primitive.M {
	var results []bson.M
	var shortList []primitive.M

	collection := databases.GetCollection(TIPS_COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	randomIdsRange := helpers.RandomNumberGenerator(len(results), count)

	for _, id := range randomIdsRange {
		shortList = append(shortList, results[id])
	}

	return shortList
}
