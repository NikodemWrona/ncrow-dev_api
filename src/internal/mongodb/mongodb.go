package mongodb

import (
	"context"
	"github.com/nikodemwrona/ncrow-dev_api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectToMongoDb() {
	user := os.Getenv("MONGODB_USER")
	password := os.Getenv("MONGODB_PASSWORD")

	mongoDbConnectionString := "mongodb://" + user + ":" + password + "@localhost:27017"

	var err error
	models.DatabaseClient, err = mongo.NewClient(options.Client().ApplyURI(mongoDbConnectionString))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = models.DatabaseClient.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
