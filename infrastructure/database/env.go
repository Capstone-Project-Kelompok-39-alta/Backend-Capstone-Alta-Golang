package database

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvDatabase() Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
	return Config{
		JWT_KEY:     os.Getenv("JWT_KEY"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		BASE_URL:    os.Getenv("BASE_URL"),
	}
}
