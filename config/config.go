package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Make sure to add this package
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
