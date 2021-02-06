package main

import (
	"github.com/gorilla/mux"
	"github.com/nikodemwrona/ncrow-dev_api/controllers"
	"github.com/nikodemwrona/ncrow-dev_api/internal/server"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	controllers.Handle(router)
	server.Start(router)
}
