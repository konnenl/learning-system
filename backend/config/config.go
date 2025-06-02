package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	JWTSecretKey string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}
	jwtSecret := getEnv("JWT_SECRET_KEY", "")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY must be set in env")
	}

	return &Config{
		ServerPort:   getEnv("SERVER_PORT", ":8080"),
		JWTSecretKey: jwtSecret,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return value
}
