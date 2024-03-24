package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DatabaseDSN() string {
	return os.Getenv("DATABASE_DSN")
}
