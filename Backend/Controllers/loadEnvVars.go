package controllers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load enviormental variables
func LoadEnvVars() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("There was an error to load envoirmental variables, %s", err)
	}
}