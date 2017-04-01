package main

import (
	"./routes"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Printf("Listening to port %s", port)
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(port, nil)
}
