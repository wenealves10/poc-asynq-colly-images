package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Variables for the database
var (
	ENV = "development"
)

// Variables for the database
var (
	REDIS_HOST     = ""
	REDIS_PORT     = ""
	REDIS_USERNAME = ""
	REDIS_PASSWORD = ""
)

// Variables for AWS S3
var (
	ACCESS_KEY_ID     = ""
	SECRET_ACCESS_KEY = ""
	REGION            = ""
	BUCKET            = ""
	ENPOINT_URL       = ""
)

func init() {
	ENV = os.Getenv("ENV")

	if ENV != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Loading environment variables for %s\n", ENV)

	// AWS S3
	ACCESS_KEY_ID = os.Getenv("ACCESS_KEY_ID")
	SECRET_ACCESS_KEY = os.Getenv("SECRET_ACCESS_KEY")
	REGION = os.Getenv("REGION")
	BUCKET = os.Getenv("BUCKET_NAME")
	ENPOINT_URL = os.Getenv("ENPOINT_URL")

	// REDIS
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_PORT = os.Getenv("REDIS_PORT")
}
