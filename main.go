package main

import (
	"log"
	"net/http"

	"api/handlers"
)

func main() {
	http.HandleFunc("/v1/ping", handlers.PingHandler)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}