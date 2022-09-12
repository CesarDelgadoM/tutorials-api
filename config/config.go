package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
