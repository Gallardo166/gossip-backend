package helper

import (
	"encoding/json"
	"gossip-backend/models"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data []*models.Post) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		WriteError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)

	if err != nil {
		WriteError(w, err)
	}
}
