package go_jwt_auth

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeErrorMessage(w http.ResponseWriter, code uint16, message string) {
	response, _ := json.Marshal(ErrorStruct{Message: message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(code))
	_, err := w.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}

type ErrorStruct struct {
	Message string `json:"message"`
}
