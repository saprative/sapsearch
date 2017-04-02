package routes

import (
	"encoding/json"
	"fmt"
	"github.com/saprative/sapsearch/models"
	"github.com/saprative/sapsearch/utils"
	"log"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	google_ch := make(chan string)
	duck_ch := make(chan string)

	Google := new(models.GoogleModel)
	Duckduckgo := new(models.DuckDuckGoModel)
	MySearch := new(models.MySearchModel)

	start := time.Now()

	google_url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyDvhG9Yg_OfscGLx9nKLZY7iR5r-yeN_xg&cx=017576662512468239146:omuauf_lfve&q=" + q
	duckduckgo_url := "http://api.duckduckgo.com/?q=" + q + "&format=json"

	go func() {
		utils.MakeRequest(google_url, Google)
		google_ch <- Google.Items[0].Snippet
	}()

	go func() {
		utils.MakeRequest(duckduckgo_url, Duckduckgo)
		duck_ch <- Duckduckgo.RelatedTopics[0].Text
	}()

	MySearch.Results.Google.URL = google_url
	MySearch.Results.Google.Text = <-google_ch
	MySearch.Results.Duckduckgo.URL = duckduckgo_url
	MySearch.Results.Duckduckgo.Text = <-duck_ch

	out, err := json.Marshal(MySearch)
	if err != nil {
		panic(err)
	}
	log.Printf(string(out))
	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(w, string(out))
}
