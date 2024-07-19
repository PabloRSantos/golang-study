package main

import (
	"go-api/app/config"
	factory "go-api/app/factories"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()
	db := config.ConnectDb()

	factory.SetupEvent(server, db)
	factory.SetupUser(server, db)

	server.Run(":8000")
}
