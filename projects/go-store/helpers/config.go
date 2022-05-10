package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}
}
