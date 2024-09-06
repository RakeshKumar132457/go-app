package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() (*Config, error) {

	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("Invalid SERVER_PORT: %v", err)
	}
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf(("Invalid DB_PORT: %v"), err)
	}

	return &Config{
		ServerPort: port,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "go_app"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}
