package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseClient *mongo.Client

type Message struct {
	Message string `json:”message” bson:”message”`
	Name    string `json:”name” bson:”name”`
}

func SaveMessage(message Message) (bool, error) {
	databaseName := "ncrow_dev"
	collectionName := "messages"

	collection := DatabaseClient.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), message)

	if err != nil {
		return false, err
	}

	return true, nil
}
