package courses

import (
	"log"
	"app/models"

	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	CollectionName = "courses"
)

type CoursesRepository {
	DbSession: sqlbuilder.Database
}

func NewCoursesRepository(session sqlbuilder.Database) *CoursesRepository {
	return CoursesRepository{session}
}

func (r *CoursesRepository) GetAllCourses() (*[]models.Course, error) {
	var courses []models.Course
	err = r.session.Collection(CollectionName).Find().All(&courses)
	if err != nil {
		return nil, err
	}
	return &courses, nil
}

func (r *CoursesRepository) Create(course models.NewCourse) error {
	r.session.Collection(CollectionName).Insert(course)
	return nil
}