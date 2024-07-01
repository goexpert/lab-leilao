package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading variables")
		return
	}
}
