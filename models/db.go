package models

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	
	"upper.io/db.v3/postgresql"

	"app/config"
)

var session sqlbuilder.Database

func NewSession() (sqlbuilder.Database, error) {
	if session == nil {
		var err error
		cfg := config.NewConfig()
		settings := postgresql.ConnectionURL{
			Host:     cfg.Get("database.host"),
			Database: cfg.Get("database.name"),
			User:     cfg.Get("database.user"),
			Password: cfg.Get("database.password"),
		}
		session, err = postgresql.Open(settings)
		if err != nil {
			log.Printf("Can not connect to database")
			return nil, err
		}
	}
	return session, nil
}
