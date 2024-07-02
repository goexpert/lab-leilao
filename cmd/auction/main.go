package main

import (
	"context"
	"log"

	"github.com/goexpert/lab-leilao/configuration/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
