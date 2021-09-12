package api

import (
	"log"
	"os"

	"github.com/amr9mohamed/mainApp/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env varibales %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// uncomment next line to seed the database if empty
	// seed.Load(server.DB)

	server.Run(":8080")
}
