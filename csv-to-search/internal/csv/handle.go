package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
)

var keywords []string

func CsvFileUpload(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		http.Error(w, "Error Read", http.StatusBadRequest)
		return
	}

	for i, row := range records {
		keyword := row[0]
		keywords = append(keywords, keyword)
		fmt.Fprintf(w, "Row %d: %v\n", i, keyword)
	}

	fmt.Println(records)

}

func SearchResult(w http.ResponseWriter, r *http.Request) {

	resultChannel := make(chan KeywordResult, len(keywords))

	for _, keyword := range keywords {
		go searchKeyword(keyword, resultChannel)
	}

	var results []KeywordResult

	for i := 0; i < len(keywords); i++ {

		res := <-resultChannel

		results = append(results, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

const SEARCH_KEY string = "-LH6yKeNphlCN2Lk"
const CSE_ID string = "15c8f4e09fb5f2"
const SEARCH_URL = "https://www.googleapis.com/customsearch/v1?key=" + SEARCH_KEY + "&cx=" + CSE_ID + "&q="
