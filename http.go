package main

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler() {
	fmt.Println("HTTP BASIC AUTH")
	http.HandleFunc("/healthcheck", check)

	http.ListenAndServe(":8090", nil)
}

func check(w http.ResponseWriter, req *http.Request) {
	result := `{"alive": true}`
	fmt.Fprintf(w, result)

}
