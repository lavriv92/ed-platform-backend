package users

import (
	"net/http"
	"encoding/json"
	"app/models"
)

func UsersCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user models.NewUser
	json.NewDecoder(r.Body).Decode(&user)
	user.SetPassword(user.Password)
	err := Create(user)
	if err != nil {
		http.Error(w, "cannot connect to database", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}