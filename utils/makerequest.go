package makerequest

import (
	"encoding/json"
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
