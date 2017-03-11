package models

import (
	"log"

	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Host:     "db:5432",
	Database: "ed_platform",
	User:     "ed_platform_user",
	Password: "ed_platform_password",
}

func NewSession() {
	session, err := postgresql.Open(settings)
	if err != nil {
		log.Printf("Can not connect to database")
	}

	defer session.Close()

	var users []User
	err = session.Collection("users").Find().All(&users)
	if err != nil {
		log.Printf("Find(): %q\n", err)
	}

	for i, user := range users {
		log.Printf("Book %d: %#v\n", i, user)
	}
}
