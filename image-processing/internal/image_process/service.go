package img_processing

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, errors []string) {
	res := ErrorResponse{
		Status: false,
		Code:   statusCode,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func sendSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := SuccessResponse{
		Status: true,
		Code:   statusCode,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}
