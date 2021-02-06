package controllers

import (
	"github.com/gorilla/mux"
	HealthCheck "github.com/nikodemwrona/ncrow-dev_api/controllers/healthCheck"
	"net/http"
)

func Handle(router *mux.Router) {
	router.HandleFunc(HealthCheck.PATH, HealthCheck.Controller).Methods(http.MethodGet)
}
