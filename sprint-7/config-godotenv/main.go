package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	log.Println(s3Bucket, secretKey)

	configEnv, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err)
	}
	log.Println(configEnv)
}
