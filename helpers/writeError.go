package helper

import (
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
