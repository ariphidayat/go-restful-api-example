package api

import (
	"fmt"
	"github.com/ariphidayat/go-restful-api-example/api/controllers"
	"github.com/ariphidayat/go-restful-api-example/api/seed"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func init() {
	// loads .env file
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
	}
}

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	} else {
		fmt.Println(".env loaded successfully.")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	seed.Load(server.DB)

	server.Run(":8080")
}