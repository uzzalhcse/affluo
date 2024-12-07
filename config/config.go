package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL       string
	RedisURL          string
	Port              int
	JWTSecret         string
	TrackingSecretKey string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	config := &Config{
		DatabaseURL:       getEnv("DATABASE_URL", "postgresql://postgres:root@localhost:5432/affluo"),
		RedisURL:          getEnv("REDIS_URL", "localhost:6379"),
		Port:              getEnvAsInt("PORT", 8080),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key"),
		TrackingSecretKey: getEnv("TRACKING_SECRET", "tracking-secret-key"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
