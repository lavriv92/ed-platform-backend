package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func FormatMessage(message string, status int) ErrorMessage {
	return ErrorMessage{message, status}
}

func RouteNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	error := FormatMessage("Route not found", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	json.NewEncoder(w).Encode(error)
}
