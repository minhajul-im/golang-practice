package router

import (
	"log"
	"net/http"
	"webserver/controller"
)

func Server() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", controller.GetUsers)
	mux.HandleFunc("POST /users/store", controller.StoreUser)
	mux.HandleFunc("PATCH /users/update/{id}", controller.UpdateUser)
	mux.HandleFunc("DELETE /users/delete/{id}", controller.DeleteUser)

	log.Println("SERVER STARTING on :8080...")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {

		log.Fatalln("SERVER START ERROR:")
	}
}
