package models

type NewCourse struct {
	AuthorId    float64 `db:"author_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
}

type Course struct {
	*NewCourse
	ID          float64 `db:"id"`
}