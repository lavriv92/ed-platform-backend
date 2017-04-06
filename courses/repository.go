package courses
import (
	"app/models"
)

func GetAllCourses (*[]models.Course, error) {
	session, err = models.NewSession()
	if err != nil {
		return nil, err
	}
	var courses []models.Course
	err := session.Collection("courses").Find().All(&courses)
	if err != nil {
		return nil, err
	}
	return &courses, nil
}