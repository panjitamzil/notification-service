package main

import (
	"log"
	"notification-service/internal/api"
	"notification-service/internal/config"
	"notification-service/internal/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	logger.Init()
	config.Load()

	router := api.SetupRouter()

	log.Println("Starting server on :8080")
	if err := router.Start(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
