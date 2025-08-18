package service

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

func GetUserList() ([]model.User, error) {
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

func SaveUserList(users []model.User) error {
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

func GetUserId(r *http.Request, pathName string) (int, error) {
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
