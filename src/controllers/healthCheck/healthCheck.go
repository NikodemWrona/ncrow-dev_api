package healthCheck

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Ok bool `json:"ok"`
}

const PATH = "/healthcheck"

const ORIGIN = "*"
const METHODS = "POST, OPTIONS"
const HEADERS = "Accept, Content-Type, Content-Length, Authorization"

func Controller(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s REQUEST : %s \n", request.Method, PATH)

	w.Header().Set("Access-Control-Allow-Origin", ORIGIN)
	w.Header().Set("Access-Control-Allow-Methods", METHODS)
	w.Header().Set("Access-Control-Allow-Headers", HEADERS)

	w.Header().Set("Content-Type", "application/json")

	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	payload := Payload{Ok: true}
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
