package auth

import (
	"encoding/json"
	"net/http"
	"app/users"
)

type TokenRequestData struct {
	Email    string `json:"email"`
	Password string `json: "password"`
}

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	var requestData TokenRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := users.FindByEmail(requestData.Email)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil || !user.ValidPassword(requestData.Password) {
		http.Error(w, "User can not found or password is incorrect", http.StatusNotFound);
		return
}
	token, err := CreateToken(user)
	if err != nil {
		http.Error(w, "Error with create token", http.StatusInternalServerError)
		return
	}
	response := AuthSuccessResponse{token, user}
	json.NewEncoder(w).Encode(response)
}
