package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprint(w, err)
	log.Print(err)
}

func WriteJsonError(w http.ResponseWriter, message string, status int) {
	var errorJson struct {
		Message string `json:"message"`
	}
	errorJson.Message = message
	jsonData, err := json.Marshal(errorJson)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}
}
