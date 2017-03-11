package models

type User struct {
	ID        uint64 `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	// CreatedAt time.Time `db:"created_time" json:"created_at"`
	// UpdatedAt time.Time `db:"updated_time" json:"updated_at"`
}
