package courses

import (
	"log"
	"app/models"
)

const (
	CollectionName = "courses"
)

func GetAllCourses() (*[]models.Course, error) {
	session, err := models.NewSession()
	if err != nil {
		log.Fatal("error connetion to database")
		return nil, err
	}
	var courses []models.Course
	err = session.Collection(CollectionName).Find().All(&courses)
	if err != nil {
		return nil, err
	}
	return &courses, nil
}

func Create(course models.NewCourse) error {
	session, err := models.NewSession()
	if err != nil {
		return err
	}
	session.Collection(CollectionName).Insert(course)
	return nil
}