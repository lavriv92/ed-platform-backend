package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := CreateToken()
	if err != nil {
		log.Printf("Error of token log")
	}
	response := AuthSuccessResponse{token}
	json.NewEncoder(w).Encode(response)
}
