package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ResponseJson(w http.ResponseWriter, code int, message string, data interface{}) {
	response := Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}

	responseJSON, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(responseJSON)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, message, nil)
}
