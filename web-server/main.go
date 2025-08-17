package main

import (
	"webserver/router"
)

func main() {

	router.Server()
	// mux := http.NewServeMux()

	// mux.Handle("/contact", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "METHOD IS THE CONTACT USE THE HANDLER")
	// }))

	// mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Matched: GET /USER")
	// })

	// log.Println("SERVER STARTING on :8080...")

	// err := http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	log.Fatalf("SERVER START ERROR: %v", err)
	// }
}
