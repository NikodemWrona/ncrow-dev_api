package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nikodemwrona/ncrow-dev_api/src/controllers"
	"github.com/nikodemwrona/ncrow-dev_api/src/internal/mongodb"
	"github.com/nikodemwrona/ncrow-dev_api/src/internal/server"
	"github.com/nikodemwrona/ncrow-dev_api/src/models"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	mongodb.ConnectToMongoDb()
	defer models.DatabaseClient.Disconnect(context.Background())

	controllers.Handle(router)
	server.Start(router)
}
