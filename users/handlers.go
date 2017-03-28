package users

import (
	"net/http"
	"encoding/json"
	"app/models"
)

func UsersCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.SetPassword(user.Password)
	Create(user)
	json.NewEncoder(w).Encode(user)
}