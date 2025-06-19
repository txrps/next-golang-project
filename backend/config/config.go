package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DatabaseURL  string
	Environment  string
	LogLevel     string
	AllowOrigins string
	JWTSecret    string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading file: %v", err)
	}

	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		AllowOrigins: getEnv("ALLOW_ORIGINS", "http://localhost:3000"),
		JWTSecret:    getEnv("JWT_Secret", "secret-jwt"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
