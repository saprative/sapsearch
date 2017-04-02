package main

import (
	"github.com/saprative/sapsearch/routes"
	"log"
	"net/http"
	"os"
)

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		log.Printf("INFO: No PORT environment variable detected.")
	}
	return ":" + port
}

func main() {
	port := GetPort()
	log.Printf("Listening to port %s", port)
	http.HandleFunc("/", routes.Handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf(err.Error())
	}
}
