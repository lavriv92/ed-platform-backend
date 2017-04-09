package courses

import (
	"log"
	"encoding/json"
	"net/http"
	"app/models"
	
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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

func CourseGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	course, err := GetCourse(vars["id"])
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	json.NewEncoder(w).Encode(course)
}

func CoursesCreateHandler(w http.ResponseWriter, r *http.Request) {
	var course models.NewCourse
	json.NewDecoder(r.Body).Decode(&course)
	userId := context.Get(r, "currentUserId")
	log.Printf("%s", userId)
	course.SetAuthorID(userId.(uint64))
	err := Create(course)
	if err != nil {
		http.Error(w, "Error create course", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(course)
}
