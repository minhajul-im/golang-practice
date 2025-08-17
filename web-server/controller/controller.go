package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userList = make([]User, 0)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	// u1 := User{
	// 	ID:    len(userList) + 1,
	// 	Name:  "Minhaj",
	// 	Email: "minhaj@gmail.com",
	// }

	// userList = append(userList, u1)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(userList)

	if err != nil {
		http.Error(w, "JSON Convert Error!", 400)
		return
	}

}

func StoreUser(w http.ResponseWriter, r *http.Request) {

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid JSON data:", 404)
		return
	}

	newUser.ID = len(userList) + 1

	userList = append(userList, newUser)

	w.Header().Set("Content-Type", "application/json")

	compileErr := json.NewEncoder(w).Encode(newUser)

	if compileErr != nil {
		http.Error(w, "Failed to encode response:", 404)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user User

	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 3 || parts[0] != "users" || parts[1] != "update" {
		http.Error(w, "PATH is not valid", 400)
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "Invalid User Id", 400)
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON data:", 404)
		return
	}

	found := false
	for i, existingUser := range userList {
		if existingUser.ID == id {
			user.ID = id
			userList[i] = user
			found = true
			break
		}
	}

	if !found {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 3 || parts[0] != "users" || parts[1] != "delete" {
		http.Error(w, "Invalid URL format. Use /users/delete/<id>", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "Invalid User Id", 400)
	}

	found := false
	for i, existingUser := range userList {
		if existingUser.ID == id {

			userList = append(userList[:i], userList[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User deleted successfully"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
