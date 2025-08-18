package server

import (
	"csv_to_search/cmd/router"
	"log"
	"net/http"
)

func Server() {

	mux := http.NewServeMux()

	router.Router(mux)

	log.Println("SERVER STARTING on :8080...")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalln("SERVER START ERROR:", err)
	}
}
