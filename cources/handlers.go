package cources

import (
	"encoding/json"
	"net/http"
	"time"
)

type Course struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"updated"`
}

type Courses []Course

func CourcesIndexHandler(w http.ResponseWriter, r *http.Request) {
	courses := Courses{
		Course{
			Name:        "Software engineering",
			Description: "Description",
		},
		Course{
			Name:        "Software engineering 2",
			Description: "Description",
		},
	}
	json.NewEncoder(w).Encode(courses)
}
