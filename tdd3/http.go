package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("HTTP BASIC AUTH")
	http.HandleFunc("/healthcheck", check)

	http.ListenAndServe(":8090", nil)
}

func check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
	io.WriteString(w, `{"alive": true}`)
}
