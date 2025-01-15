package helper

import (
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	log.Fatalf("Error: %s", err)
}
