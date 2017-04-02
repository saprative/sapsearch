package router

import (
	"../models"
	"../utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")
	//a := make(chan interface{})
	//b := make(chan interface{})

	Google := new(models.GoogleModel)
	Duckduckgo := new(models.DuckDuckGoModel)
	MySearch := new(models.MySearchModel)

	google_url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyDvhG9Yg_OfscGLx9nKLZY7iR5r-yeN_xg&cx=017576662512468239146:omuauf_lfve&q=" + q
	duckduckgo_url := "http://api.duckduckgo.com/?q=" + q + "&format=json"

	makerequest.MakeRequest(google_url, Google)

	//if err_google != nil {
	//log.Printf(err_google.Error())
	//}

	makerequest.MakeRequest(duckduckgo_url, Duckduckgo)

	//if err_duckduckgo != nil {
	//log.Printf(err_duckduckgo.Error())
	//}

	MySearch.Results.Google.URL = google_url
	MySearch.Results.Google.Text = Google.Items[0].Snippet
	MySearch.Results.Duckduckgo.URL = duckduckgo_url
	MySearch.Results.Duckduckgo.Text = Duckduckgo.RelatedTopics[0].Text

	out, err := json.Marshal(MySearch)
	if err != nil {
		panic(err)
	}

	log.Printf(string(out))
	fmt.Fprintf(w, string(out))
}
