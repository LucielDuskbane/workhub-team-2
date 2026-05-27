package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		// ganti log.Fatal() ke log.Println() biar
		// nggak crash waktu saat file .env kagak ada
		// waktu docker build
		log.Println("No .env file found, using injected environment variables")
	}
}
