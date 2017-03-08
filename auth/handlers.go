package auth

import (
	"encoding/json"
	"net/http"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	response := AuthSuccessResponse{ "some token" }
	json.NewEncoder(w).Encode(response)
}