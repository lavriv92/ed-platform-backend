package cources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/models"
)

func CourcesIndexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := models.NewSession()
	if err != nil {
		log.Printf("Format err")
	}
	var users []models.User
	err = session.Collection("users").Find().All(&users)
	if err != nil {
		log.Printf("Error select users")
	}

	for _, user := range users {
		fmt.Printf("%s .\n", user.FirstName)
	}

	json.NewEncoder(w).Encode(users)
}
