package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Msg struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a := Msg{1, "teste", "teste"}
	fmt.Print(a)
	json.NewEncoder(w).Encode(a)

}

func obter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/get", obter)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func db() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
