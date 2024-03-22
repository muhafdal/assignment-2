package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Dialect    string
	Port       string
}

func LoadAppConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error while load .env file")
	}
}

func GetAppConfig() appConfig {
	return appConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Dialect:    os.Getenv("DIALECT"),
		Port:       os.Getenv("PORT"),
	}
}
