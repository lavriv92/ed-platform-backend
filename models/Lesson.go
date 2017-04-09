package models

type NewLesson struct {
	CourseId    uint64 `db:"course_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type Lesson struct {
	*NewLesson,
	Id uint64 `db:"id"`
}

func (l *NewLesson)SetCourse(courseId uint64) {
	l.CourseId = courseId
}