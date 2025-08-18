package controller

import (
	"encoding/json"
	"net/http"
	"webserver/model"
	"webserver/service"
	"webserver/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := service.GetUserList()

	if err != nil {
		var errors []string
		errors = append(errors, "Server Error")
		utils.SendErrorRes(w, 500, errors)
		return
	}

	utils.SendSuccessRes(w, http.StatusOK, users)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil || r.ContentLength == 0 {
		var errors []string
		errors = append(errors, "Empty Request body")
		utils.SendErrorRes(w, http.StatusBadRequest, errors)
		return
	}

	var newUser model.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		var errors []string
		errors = append(errors, "Invalid JSON data")
		utils.SendErrorRes(w, http.StatusBadRequest, errors)
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

		utils.SendErrorRes(w, http.StatusBadRequest, errs)
		return

	}

	users, err := service.GetUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		utils.SendErrorRes(w, http.StatusServiceUnavailable, errs)
		return
	}

	for _, user := range users {
		if user.Email == newUser.Email {
			var errors []string
			errors = append(errors, "User Already Exits")
			utils.SendErrorRes(w, http.StatusBadRequest, errors)
			return
		}
	}

	newUser.ID = len(users) + 1

	users = append(users, newUser)

	err = service.SaveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to save user")
		utils.SendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	utils.SendSuccessRes(w, http.StatusCreated, newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	id, err := service.GetUserId(r, "update")

	if err != nil {
		utils.SendErrorRes(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		var errors []string
		errors = append(errors, "Invalid JSON Body")
		utils.SendErrorRes(w, http.StatusBadRequest, errors)
		return
	}

	users, err := service.GetUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		utils.SendErrorRes(w, http.StatusServiceUnavailable, errs)
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
		utils.SendErrorRes(w, http.StatusBadRequest, errs)
		return
	}

	err = service.SaveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to save user")
		utils.SendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	utils.SendSuccessRes(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := service.GetUserId(r, "delete")

	if err != nil {
		utils.SendErrorRes(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	users, err := service.GetUserList()

	if err != nil {
		var errs []string
		errs = append(errs, "Server Error")
		utils.SendErrorRes(w, http.StatusServiceUnavailable, errs)
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
		utils.SendErrorRes(w, http.StatusBadRequest, errs)
		return
	}

	err = service.SaveUserList(users)
	if err != nil {
		var errors []string
		errors = append(errors, "Failed to delete user")
		utils.SendErrorRes(w, http.StatusInternalServerError, errors)
		return
	}

	utils.SendSuccessRes(w, http.StatusOK, "User has been deleted")
}
