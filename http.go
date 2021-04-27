package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HTTP BASIC AUTH")
	http.HandleFunc("/healthcheck", check)

	http.ListenAndServe(":8090", nil)
}

func check(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")

}
