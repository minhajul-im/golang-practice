package router

import (
	"fmt"
	"net/http"
)

func Router(mux *http.ServeMux) {

	mux.HandleFunc("POST /auth/signin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello handFunc")
	})

	mux.HandleFunc("GET /auth/signout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello handFunc")
	})

}
