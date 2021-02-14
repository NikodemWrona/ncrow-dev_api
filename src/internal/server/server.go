package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const PORT = 8081

func Start(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))

	serverAddress := ":" + strconv.Itoa(PORT)
	fmt.Println("Server started")
	fmt.Println("PORT : ", PORT)

	err := http.ListenAndServe(serverAddress, router)

	if err != nil {
		log.Fatal(err)
	}
}
