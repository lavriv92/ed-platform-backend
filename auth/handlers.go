package auth

import (
	"encoding/json"
	"net/http"
)

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	response := AuthSuccessResponse{"some token"}
	json.NewEncoder(w).Encode(response)
}
