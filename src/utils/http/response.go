package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ErrorResponse(rw http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Status:     "error",
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}

	sendResponse(rw, statusCode, response)
}

func SuccessResponse(rw http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := Response{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	sendResponse(rw, statusCode, response)
}

func sendResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
