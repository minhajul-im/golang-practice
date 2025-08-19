package csv

type KeywordResult struct {
	Keyword string `json:"keyword"`
	Result  string `json:"result"`
}

type GoogleSearchItem struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

type GoogleSearchResponse struct {
	Items []GoogleSearchItem `json:"items"`
}
