package database

// env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvDatabase() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		panic("error loading .env file")
	}
	return Config{
		JWT_KEY:     os.Getenv("JWT_KEY"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
	}
}
