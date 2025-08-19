package router

import (
	"csv_to_search/internal/csv"
	"csv_to_search/internal/user"
	"net/http"
)

func Router(mux *http.ServeMux) {

	mux.HandleFunc("POST /auth/signin", user.UserSignin)

	mux.HandleFunc("GET /auth/signout", user.UserSignout)

	mux.HandleFunc("POST /files/store", csv.CsvFileUpload)

	mux.HandleFunc("GET /search-result", csv.SearchResult)

}
