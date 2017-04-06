package courses

import (
	"log"
	"encoding/json"
	"net/http"
)

func CoursesIndexHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := GetAllCourses()
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Courses not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	json.NewEncoder(w).Encode(courses)
}
