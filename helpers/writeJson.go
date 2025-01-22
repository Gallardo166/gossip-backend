package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	fmt.Print(data)
	jsonData, err := json.Marshal(data)
	fmt.Print(jsonData)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)

	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}
}
