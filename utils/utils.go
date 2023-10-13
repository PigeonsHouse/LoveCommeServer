package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiPort    string
	DbHost     string
	DbUser     string
	DbPass     string
	DbName     string
	DbPort     string
	SigningKey []byte
)

func LoadEnv() {
	godotenv.Load(".env")
	ApiPort = getEnv("PORT", "8000")
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
