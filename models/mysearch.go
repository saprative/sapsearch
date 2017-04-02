package models

type MySearchModel struct {
	Query   string `json:"query"`
	Results struct {
		Google struct {
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"google"`
		Twitter struct {
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"twitter"`
		Duckduckgo struct {
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"duckduckgo"`
	} `json:"results"`
}
