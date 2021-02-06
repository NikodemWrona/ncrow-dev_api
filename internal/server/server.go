package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080

func Start(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))

	serverAddress := ":" + strconv.Itoa(PORT)
	err := http.ListenAndServe(serverAddress, router)

	if err != nil {
		log.Fatal(err)
	}
}
