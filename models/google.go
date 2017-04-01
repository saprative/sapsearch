package models

type GoogleModel struct {
	Context struct {
		Facets []struct {
			Num0 struct {
				Anchor      string `json:"anchor"`
				Label       string `json:"label"`
				LabelWithOp string `json:"label_with_op"`
			} `json:"0"`
		} `json:"facets"`
		Title string `json:"title"`
	} `json:"context"`
	Items []struct {
		CacheID          string `json:"cacheId"`
		DisplayLink      string `json:"displayLink"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		HTMLSnippet      string `json:"htmlSnippet"`
		HTMLTitle        string `json:"htmlTitle"`
		Kind             string `json:"kind"`
		Link             string `json:"link"`
		Pagemap          struct {
			CseImage []struct {
				Src string `json:"src"`
			} `json:"cse_image"`
			CseThumbnail []struct {
				Height string `json:"height"`
				Src    string `json:"src"`
				Width  string `json:"width"`
			} `json:"cse_thumbnail"`
			Metatags []struct {
				OgDescription string `json:"og:description"`
				OgTitle       string `json:"og:title"`
				OgType        string `json:"og:type"`
				TwitterCard   string `json:"twitter:card"`
				Viewport      string `json:"viewport"`
			} `json:"metatags"`
		} `json:"pagemap,omitempty"`
		Snippet    string `json:"snippet"`
		Title      string `json:"title"`
		FileFormat string `json:"fileFormat,omitempty"`
		Mime       string `json:"mime,omitempty"`
	} `json:"items"`
	Kind    string `json:"kind"`
	Queries struct {
		NextPage []struct {
			Count          int    `json:"count"`
			Cx             string `json:"cx"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			SearchTerms    string `json:"searchTerms"`
			StartIndex     int    `json:"startIndex"`
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
		} `json:"nextPage"`
		Request []struct {
			Count          int    `json:"count"`
			Cx             string `json:"cx"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			SearchTerms    string `json:"searchTerms"`
			StartIndex     int    `json:"startIndex"`
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
		} `json:"request"`
	} `json:"queries"`
	SearchInformation struct {
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
		SearchTime            float64 `json:"searchTime"`
		TotalResults          string  `json:"totalResults"`
	} `json:"searchInformation"`
	URL struct {
		Template string `json:"template"`
		Type     string `json:"type"`
	} `json:"url"`
}
