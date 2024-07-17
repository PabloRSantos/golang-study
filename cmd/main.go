package main

import (
	"go-api/app/config"
	factory "go-api/app/factories"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db := config.ConnectDb()

	factory.SetupEvent(server, db)
	factory.SetupUser(server, db)

	server.Run(":8000")
}
