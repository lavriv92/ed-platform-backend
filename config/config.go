package config

import (
	"log"
	"path/filepath"
	"github.com/moraes/config"
)

type Config struct {
	cfg *config.Config
}

const fileByteSize = 100

var conf *Config

func NewConfig() *Config {
	if conf == nil {
		conf = &Config{nil}
	}
	return conf
}

func (c *Config) Load(configFile string) {
	if c.cfg == nil {
		configPath := filepath.Join(filepath.Base("."), "config.yml")
		cfg, err := config.ParseYamlFile(configPath)
		if err != nil {
			log.Fatal("Error parsing config: %s", err.Error())
		}
		c.cfg = cfg
	}
}

func (c *Config)Get(key string) string {
	value, err := c.cfg.String(key)
	if err != nil {
		log.Fatal("Can not load value by key: %s ", err.Error())
	}
	return value
}