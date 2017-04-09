package lessons

import (
	"log"
	"app/models"
)

const (
	CollectionName = "lessons"
)

func GetLessonByCourseId(courseId string) ([]*models.Lesson, error) {
	session, err := models.NewSession()
	if err != nil {
		log.Printf("error connect to database")
		return nil, err
	}
	var lesson []models.Lesson
	err = session.Collection(CollectionName).Find().Where("course_id = ?", courseId).All(&lesson)
	if err != nil {
		log.Printf("Cannot found any lessons")
		return nil, err
	}
	return &lessons, nil
}