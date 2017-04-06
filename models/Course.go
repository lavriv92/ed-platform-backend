package models

type NewCourse struct {
	AuthorID    uint64 `db:"author_id" json:"author_id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
}

type Course struct {
	*NewCourse
	ID          uint64 `db:"id" json:"id"`
}

func (c *NewCourse) SetAuthorID(id uint64) {
	c.AuthorID = id
}