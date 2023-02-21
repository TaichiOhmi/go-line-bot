package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprintf(w, msg)
}
