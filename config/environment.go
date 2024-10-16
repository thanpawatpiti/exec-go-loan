package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppMode string
	AppName string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// Jwt
	JwtSecretKey string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}

	// App
	AppMode = os.Getenv("APP_MODE")
	AppName = os.Getenv("APP_NAME")

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")

	// Jwt
	JwtSecretKey = os.Getenv("JWT_SECRET")
}
