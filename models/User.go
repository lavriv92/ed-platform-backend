package models

import (
	"app/shared"
)

type PasswordableUser interface {
	SetPassword(password string)
	ValidPassword(password string) bool
}

type NewUser struct {
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

type User struct {
	*NewUser
	ID        uint64 `db:"id" json:"id"`
}

func (u *NewUser) SetPassword(password string) {
	u.Password = shared.EncodePassword(password)
}

func (u *NewUser) ValidPassword(password string) bool {
	return u.Password == shared.EncodePassword(password)
}