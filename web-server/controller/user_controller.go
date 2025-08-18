package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"webserver/model"
)

func getUserList() ([]model.User, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	jsonPath := filepath.Join(wd, "db", "db.json")

	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var users []model.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func saveUserList(users []model.User) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	jsonPath := filepath.Join(wd, "db", "db.json")

	file, err := os.OpenFile(jsonPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(users)
}

func getUserId(r *http.Request, pathName string) (int, error) {
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 3 || parts[0] != "users" || parts[1] != pathName {
		return 0, fmt.Errorf("PATH is not valid")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("Invalid User ID")
	}

	return id, nil
}

func sendErrorRes(w http.ResponseWriter, statusCode int, errors []string) {
	res := model.ErrorResponse{
		Status: false,
		Code:   statusCode,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func sendSuccessRes(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := model.SuccessResponse{
		Status: true,
		Code:   statusCode,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := getUserList()

	if err != nil {
		var errors []string
		errors = append(errors, "Server Error")
		sendErrorRes(w, 500, errors)
		return
	}

	sendSuccessRes(w, http.StatusOK, users)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil || r.ContentLength == 0 {
		var errors []string
		errors = append(errors, "Empty Request body")
		sendErrorRes(w, http.StatusBadRequest, errors)
		return
	}

	var newUser model.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		var errors []string
		errors = append(errors, "Invalid JSON data")
		sendErrorRes(w, http.StatusBadRequest, errors)
		return
	}

	if newUser.Email == "" || newUser.Name == "" {
		var errs []string

		if newUser.Name == "" {
			errs = append(errs, "name is required")
		}
		if newUser.Email == "" {
			errs = append(errs, "email is required")
		}

		sendErrorRes(w, http.StatusBadRequest, errs)
		return

	}

	users, err := getUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		sendErrorRes(w, http.StatusServiceUnavailable, errs)
		return
	}

	for _, user := range users {
		if user.Email == newUser.Email {
			var errors []string
			errors = append(errors, "User Already Exits")
			sendErrorRes(w, http.StatusBadRequest, errors)
			return
		}
	}

	newUser.ID = len(users) + 1

	users = append(users, newUser)

	err = saveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to save user")
		sendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	sendSuccessRes(w, http.StatusCreated, newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	id, err := getUserId(r, "update")

	if err != nil {
		sendErrorRes(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		var errors []string
		errors = append(errors, "Invalid JSON Body")
		sendErrorRes(w, http.StatusBadRequest, errors)
		return
	}

	users, err := getUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		sendErrorRes(w, http.StatusServiceUnavailable, errs)
		return
	}

	found := false
	for i, existingUser := range users {
		if existingUser.ID == id {
			user.ID = id
			users[i].Email = user.Email
			users[i].Name = user.Name
			found = true
			break
		}
	}

	if !found {
		var errs []string
		errs = append(errs, "User is not found!")
		sendErrorRes(w, http.StatusBadRequest, errs)
		return
	}

	err = saveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to save user")
		sendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	sendSuccessRes(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := getUserId(r, "delete")

	if err != nil {
		sendErrorRes(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	users, err := getUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		sendErrorRes(w, http.StatusServiceUnavailable, errs)
		return
	}

	found := false
	for i, existingUser := range users {
		if existingUser.ID == id {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		var errs []string
		errs = append(errs, "User is not found!")
		sendErrorRes(w, http.StatusBadRequest, errs)
		return
	}

	err = saveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to delete user")
		sendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	sendSuccessRes(w, http.StatusOK, "User has been deleted")
}
