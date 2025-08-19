package csv

import (
	"io"
	"net/http"
)

func searchKeyword(keyword string, results chan<- KeywordResult) {
	url := SEARCH_URL + keyword

	res, err := http.Get(url)
	if err != nil {
		results <- KeywordResult{Keyword: keyword, Result: "Error: " + err.Error()}
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		results <- KeywordResult{Keyword: keyword, Result: "Error: " + err.Error()}
		return
	}

	resData := string(body)

	results <- KeywordResult{Keyword: keyword, Result: resData}

}
