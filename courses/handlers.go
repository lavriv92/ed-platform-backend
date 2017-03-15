package courses

import (
	"encoding/json"
	"log"
	"net/http"

	"app/models"
)

func CoursesIndexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := models.NewSession()
	if err != nil {
		log.Printf("error creating session")
	}
	var users []models.User
	err = session.Collection("users").Find().All(&users)
	if err != nil {
		log.Printf("Error select users")
	}

	json.NewEncoder(w).Encode(users)
}
