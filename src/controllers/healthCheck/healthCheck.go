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

func Controller(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s REQUEST : %s \n", request.Method, PATH)

	payload := Payload{Ok: true}
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
