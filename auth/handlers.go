package auth

import (
	"encoding/json"
	"net/http"
	"app/users"
)

type RequestData struct {
	Email    string `json:"email"`
	Password string `json: "password"`
}

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestData RequestData
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	user, err := users.FindByEmail(requestData.Email)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		http.Error(w, "User can not found", http.StatusNotFound);
	} else {
		token, err := CreateToken(user.ID)
		if err != nil {
			http.Error(w, "Error with create token", http.StatusInternalServerError)
		}
		response := AuthSuccessResponse{token}
		json.NewEncoder(w).Encode(response)
	}
}
