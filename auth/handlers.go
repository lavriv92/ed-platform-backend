package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"app/users"
	"app/errors"
)

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := users.FindByEmail(vars["email"])
	if err != nil {
		response := errors.FormatMessage("User can not found", http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	} else {
		token, err := CreateToken(user.ID)
		if err != nil {
			log.Printf("Error of token log")
		}
		response := AuthSuccessResponse{token}
		json.NewEncoder(w).Encode(response)
	}
}
