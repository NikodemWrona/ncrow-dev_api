package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseClient *mongo.Client

type Message struct {
	Message string `json:”message” bson:”message”`
	Name    string `json:”name” bson:”name”`
}

func SaveMessage(message Message) (bool, error) {
	fmt.Println("SAVE MESSAGE")
	collection := DatabaseClient.Database("ncrow_dev").Collection("messages")
	_, err := collection.InsertOne(context.Background(), message)

	if err != nil {
		return false, err
	}

	return true, nil
}
