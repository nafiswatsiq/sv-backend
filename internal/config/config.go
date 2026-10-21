package config

import "os"

type Config struct {
	DBConnection string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	ServerPort   string
}

func Load() Config {
	return Config{
		DBConnection: getEnv("DB_CONNECTION", getEnv("DB_CONNECTON", "mysql")),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "3306"),
		DBUser:       getEnv("DB_USER", "root"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "sharing_vision"),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
