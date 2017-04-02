package makerequest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func MakeRequest(url string, structure interface{}) error {

	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Get(url)

	if err != nil {
		log.Printf(err.Error())
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(structure)

}
