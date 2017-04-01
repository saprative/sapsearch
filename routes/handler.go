package handler

import (
	"../models"
	"../utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	Google := new(models.GoogleModel)
	Duckduckgo := new(models.DuckDuckGoModel)
	MySearch := new(models.MySearchModel)
	q := r.URL.Query().Get("q")
	google_url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyDvhG9Yg_OfscGLx9nKLZY7iR5r-yeN_xg&cx=017576662512468239146:omuauf_lfve&q=" + q
	duckduckgo_url := "http://api.duckduckgo.com/?q=" + q + "&format=json"
	makerequest.MakeRequest(google_url, Google)
	makerequest.MakeRequest(duckduckgo_url, Duckduckgo)

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
