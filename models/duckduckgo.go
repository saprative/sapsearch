package models

type DuckDuckGoModel struct {
	DefinitionSource string `json:"DefinitionSource"`
	Heading          string `json:"Heading"`
	ImageWidth       int    `json:"ImageWidth"`
	RelatedTopics    []struct {
		Result string `json:"Result,omitempty"`
		Icon   struct {
			URL    string `json:"URL"`
			Height string `json:"Height"`
			Width  string `json:"Width"`
		} `json:"Icon,omitempty"`
		FirstURL string `json:"FirstURL,omitempty"`
		Text     string `json:"Text,omitempty"`
		Topics   []struct {
			Result string `json:"Result"`
			Icon   struct {
				URL    string `json:"URL"`
				Height string `json:"Height"`
				Width  string `json:"Width"`
			} `json:"Icon"`
			FirstURL string `json:"FirstURL"`
			Text     string `json:"Text"`
		} `json:"Topics,omitempty"`
		Name string `json:"Name,omitempty"`
	} `json:"RelatedTopics"`
	Entity string `json:"Entity"`
	Meta   struct {
		Maintainer struct {
			Github string `json:"github"`
		} `json:"maintainer"`
		PerlModule      string      `json:"perl_module"`
		Status          string      `json:"status"`
		ProductionState string      `json:"production_state"`
		DevDate         interface{} `json:"dev_date"`
		JsCallbackName  string      `json:"js_callback_name"`
		SignalFrom      string      `json:"signal_from"`
		LiveDate        interface{} `json:"live_date"`
		SrcID           int         `json:"src_id"`
		SrcOptions      struct {
			SkipEnd           string `json:"skip_end"`
			SkipAbstract      int    `json:"skip_abstract"`
			SkipQr            string `json:"skip_qr"`
			Language          string `json:"language"`
			SkipIcon          int    `json:"skip_icon"`
			SkipImageName     int    `json:"skip_image_name"`
			Directory         string `json:"directory"`
			MinAbstractLength string `json:"min_abstract_length"`
			SkipAbstractParen int    `json:"skip_abstract_paren"`
			IsWikipedia       int    `json:"is_wikipedia"`
			SourceSkip        string `json:"source_skip"`
			IsFanon           int    `json:"is_fanon"`
			IsMediawiki       int    `json:"is_mediawiki"`
			SrcInfo           string `json:"src_info"`
		} `json:"src_options"`
		Repo      string `json:"repo"`
		Developer []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"developer"`
		Tab             string      `json:"tab"`
		Producer        interface{} `json:"producer"`
		Unsafe          int         `json:"unsafe"`
		ID              string      `json:"id"`
		DevMilestone    string      `json:"dev_milestone"`
		Topic           []string    `json:"topic"`
		Name            string      `json:"name"`
		Attribution     interface{} `json:"attribution"`
		CreatedDate     interface{} `json:"created_date"`
		ExampleQuery    string      `json:"example_query"`
		Description     string      `json:"description"`
		IsStackexchange interface{} `json:"is_stackexchange"`
		Designer        interface{} `json:"designer"`
		SrcDomain       string      `json:"src_domain"`
		SrcName         string      `json:"src_name"`
		Blockgroup      interface{} `json:"blockgroup"`
		SrcURL          interface{} `json:"src_url"`
	} `json:"meta"`
	Type           string        `json:"Type"`
	Redirect       string        `json:"Redirect"`
	DefinitionURL  string        `json:"DefinitionURL"`
	AbstractURL    string        `json:"AbstractURL"`
	Definition     string        `json:"Definition"`
	AbstractSource string        `json:"AbstractSource"`
	Infobox        string        `json:"Infobox"`
	Image          string        `json:"Image"`
	ImageIsLogo    int           `json:"ImageIsLogo"`
	Abstract       string        `json:"Abstract"`
	AbstractText   string        `json:"AbstractText"`
	AnswerType     string        `json:"AnswerType"`
	ImageHeight    int           `json:"ImageHeight"`
	Answer         string        `json:"Answer"`
	Results        []interface{} `json:"Results"`
}
