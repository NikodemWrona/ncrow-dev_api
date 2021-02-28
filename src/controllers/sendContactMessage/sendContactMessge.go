package sendContactMessage

import (
	"encoding/json"
	"fmt"
	"github.com/nikodemwrona/ncrow-dev_api/src/models"
	"net/http"
)

type Payload struct {
	Ok bool `json:"ok"`
}

type Body struct {
	Message string
	Name    string
}

const PATH = "/contact/message"

const ORIGIN = "*"
const METHODS = "POST, OPTIONS"
const HEADERS = "Accept, Content-Type, Content-Length, Authorization"

func Controller(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s REQUEST : %s \n", request.Method, PATH)

	w.Header().Set("Access-Control-Allow-Origin", ORIGIN)
	w.Header().Set("Access-Control-Allow-Methods", METHODS)
	w.Header().Set("Access-Control-Allow-Headers", HEADERS)

	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var body Body
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Error : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message := models.Message{Message: body.Message, Name: body.Name}
	ok, err := models.SaveMessage(message)

	if ok != true || err != nil {
		fmt.Println("Error : ", err, "Ok : ", ok)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
