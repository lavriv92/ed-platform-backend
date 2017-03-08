package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	HOST        = "example"
	DB_USER     = "postgres"
	DB_NAME     = "postgres"
	DB_PASSWORD = nil
)

func NewConnection() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", HOST, DB_USER, DB_NAME))
	defer db.Close()
}
