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

func GetCourse(courseId string) (*models.Course, error) {
	session, err := models.NewSession()
	if err != nil {
		log.Fatal("Error connection database")
		return nil, err
	}
	var course models.Course
	err = session.Collection(CollectionName).Find().Where("id = ?", courseId).One(&course)
	if err != nil {
		return nil, err
	} 
	return &course, nil
}

func Create(course models.NewCourse) error {
	session, err := models.NewSession()
	if err != nil {
		return err
	}
	session.Collection(CollectionName).Insert(course)
	return nil
}