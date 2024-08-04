package main

import (
	controllers "blog-backend/Controllers"
	"log"

	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server := &controllers.Server{}
	server.InitDB()
	server.InitRouter()
	server.RunServer()
}