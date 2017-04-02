package main

import (
	"./routes"
	"log"
	"net/http"
)

func main() {
	port := ":8001"
	log.Printf("Listening to port %s", port)
	http.HandleFunc("/", routes.Handler)
	http.ListenAndServe(port, nil)
}
