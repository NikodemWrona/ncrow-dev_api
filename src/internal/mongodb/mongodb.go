package mongodb

import (
	"context"
	"fmt"
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
	port := os.Getenv("MONGODB_PORT")

	mongoDbConnectionString := "mongodb://" + user + ":" + password + "@localhost:" + port

	fmt.Println("STRING : ", mongoDbConnectionString)

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
