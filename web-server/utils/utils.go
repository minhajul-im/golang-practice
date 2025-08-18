package utils

import (
	"encoding/json"
	"net/http"
	"webserver/model"
)

func SendErrorRes(w http.ResponseWriter, statusCode int, errors []string) {
	res := model.ErrorResponse{
		Status: false,
		Code:   statusCode,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func SendSuccessRes(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := model.SuccessResponse{
		Status: true,
		Code:   statusCode,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}
