package router

import (
	"csv_to_search/internal/user"
	"net/http"
)

func Router(mux *http.ServeMux) {

	mux.HandleFunc("POST /auth/signin", user.UserSignin)

	mux.HandleFunc("GET /auth/signout", user.UserSignout)

}
