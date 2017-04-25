package models

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	
	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Host:     "db:5432",
	Database: "ed_platform",
	User:     "ed_platform_user",
	Password: "ed_platform_password",
}

var session sqlbuilder.Database

func NewSession() (sqlbuilder.Database, error) {
	if session == nil {
		var err error
		session, err = postgresql.Open(settings)
		if err != nil {
			log.Printf("Can not connect to database")
			return nil, err
		}
	}
	return session, nil
}
