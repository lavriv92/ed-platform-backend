package config

struct Config {
	configFile: ''
}

func NewConfig(filePath string) {
	return Config{filePath}
}

func FromString(key string) string {
	
}