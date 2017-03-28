package models

import (
	"app/shared"
)

type PasswordableUser interface {
	SetPassword(password string)
	ValidPassword(password string) bool
}

type User struct {
	ID        uint64 `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func (u *User) SetPassword(password string) {
	u.Password = shared.EncodePassword(password)
}

func (u *User) ValidPassword(password string) bool {
	return u.Password == shared.EncodePassword(password)
}