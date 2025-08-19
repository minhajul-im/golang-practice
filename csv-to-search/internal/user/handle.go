package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var user = User{
	Username: "username",
	Password: "123123",
}

var token string = ""

func UserSignin(w http.ResponseWriter, r *http.Request) {

	var input User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		SendErrorRes(w, http.StatusBadRequest, []string{"Invalid request!"})
		return
	}

	if input.Username == user.Username && input.Password == user.Password {
		token = "demo-token-" + time.Now().Format("150405")

		data := SigninToken{
			Message: "Signin successful",
			Token:   token,
		}

		SendSuccessRes(w, http.StatusOK, data)

		fmt.Println(token)
		return
	}

	SendErrorRes(w, http.StatusUnauthorized, []string{"Invalid credentials!"})

}

func UserSignout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		SendErrorRes(w, http.StatusUnauthorized, []string{"Missing token!"})
		return
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {

		SendErrorRes(w, http.StatusUnauthorized, []string{"Invalid format!"})
		return
	}

	if parts[1] == token {
		fmt.Println(token)
		token = ""
		data := Signout{
			Message: "Signout successful!",
		}
		SendSuccessRes(w, http.StatusOK, data)
		return
	}

	SendErrorRes(w, http.StatusUnauthorized, []string{"Unauthorized!"})

}
