package helper

import (
	"encoding/json"
	"gossip-backend/models"
	"net/http"
)

func WriteJson[T models.Post | []*models.PostPreview | []*models.Category | []*models.Comment](w http.ResponseWriter, data T) {
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
