package config

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {}

func NewConfig() Config {
	return Config{}
}

func (config *Config) Load() {
	filePath := filepath.Join(filepath.Base(""), "config.yml")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error load loading config from %s", filePath)
	}
	fileData := make([]byte, 100)
	n, err := file.Read(fileData)
	if err != nil {
		log.Fatal("Error parsing config")
	}
	log.Printf("%d / %s", n, string(fileData))
}