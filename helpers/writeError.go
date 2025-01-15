package helper

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	fmt.Println(err)
}
