package routes

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/saprative/sapsearch/models"
	"github.com/saprative/sapsearch/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	google_ch := make(chan string)
	duck_ch := make(chan string)
	twitter_ch := make(chan string)
	mysearch_ch := make(chan *models.MySearchModel)

	start := time.Now()

	google_url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyDvhG9Yg_OfscGLx9nKLZY7iR5r-yeN_xg&cx=017576662512468239146:omuauf_lfve&q=" + q
	duckduckgo_url := "http://api.duckduckgo.com/?q=" + q + "&format=json"
	twitter_url := "https://api.twitter.com/1.1/search/tweets.json?q=" + q

	go func() {
		Google := new(models.GoogleModel)
		utils.MakeRequest(google_url, Google)
		google_ch <- Google.Items[0].Snippet
	}()

	go func() {
		Duckduckgo := new(models.DuckDuckGoModel)
		utils.MakeRequest(duckduckgo_url, Duckduckgo)
		duck_ch <- Duckduckgo.RelatedTopics[0].Text
	}()

	go func() {
		anaconda.SetConsumerKey(os.Getenv("ConsumerKey"))
		anaconda.SetConsumerSecret(os.Getenv("ConsumerSecret"))
		api := anaconda.NewTwitterApi(os.Getenv("AccessToken"), os.Getenv("AccessTokenSecret"))
		searchResult, _ := api.GetSearch(q, nil)
		tweet := searchResult.Statuses

		twitter_ch <- tweet[0].Text
	}()

	go func() {
		MySearch := new(models.MySearchModel)
		MySearch.Results.Google.URL = google_url
		MySearch.Results.Google.Text = <-google_ch
		MySearch.Results.Duckduckgo.URL = duckduckgo_url
		MySearch.Results.Duckduckgo.Text = <-duck_ch
		MySearch.Results.Twitter.URL = twitter_url
		MySearch.Results.Twitter.Text = <-twitter_ch

		mysearch_ch <- MySearch
	}()

	out, err := json.Marshal(<-mysearch_ch)
	if err != nil {
		panic(err)
	}
	log.Printf(string(out))
	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(w, string(out))
}
