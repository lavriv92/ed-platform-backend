package config

import (
	"log"
	"os"
	"path/filepath"
	"encoding/json"

	"upper.io/db.v3/lib/sqlbuilder"
)

type Config struct {
	database: sqlbuilder.Database
}

var config = &Config{}

func (config *Config) Load() {
	filePath = filepath.Join(filepath.Base(""), "config.json")
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error load loading config from %s", filepath)
	}
	var fileData []byte 
	_, err := file.Read(fileData)
	if err != nil {
		log.Fatal("Error parsing config")
	}
	log.Printf("%s", fileData)
}