package main

import (
	"./models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func MakeRequest(url string, structure interface{}) error {

	myclient := &http.Client{Timeout: 10 * time.Second}

	res, err := myclient.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(structure)

}

func handler(w http.ResponseWriter, r *http.Request) {
	Google := new(google.GoogleModel)
	google_url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyDvhG9Yg_OfscGLx9nKLZY7iR5r-yeN_xg&cx=017576662512468239146:omuauf_lfve&q=sex"
	MakeRequest(google_url, Google)
	out, err := json.Marshal(Google)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(out))
}

func main() {
	port := ":8080"
	log.Printf("Listening to port %s", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
