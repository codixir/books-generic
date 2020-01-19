package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorBody struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func RespondWithError(err error, statusCode int, w http.ResponseWriter) string {
	body := ErrorBody{
		Error:  err.Error(),
		Status: statusCode,
	}

	content, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
	return string(content)
}

func RespondWithSuccess(response interface{}, w http.ResponseWriter) string {
	content, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return string(content)
}
