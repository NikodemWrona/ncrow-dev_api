package controllers

import (
	"github.com/gorilla/mux"
	HealthCheck "github.com/nikodemwrona/ncrow-dev_api/src/controllers/healthCheck"
	SendContactInfo "github.com/nikodemwrona/ncrow-dev_api/src/controllers/sendContactMessage"
	"net/http"
)

func Handle(router *mux.Router) {
	router.HandleFunc(HealthCheck.PATH, HealthCheck.Controller).Methods(http.MethodGet)
	router.HandleFunc(SendContactInfo.PATH, SendContactInfo.Controller).Methods(http.MethodPost, http.MethodOptions)
}
