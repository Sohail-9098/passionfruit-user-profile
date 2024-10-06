package main

import (
	"context"
	"log"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/config"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/routes"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to database
	client := database.Connect(config.MongoDbUri)
	defer client.Disconnect(context.TODO())
	// setup server
	server := gin.Default()
	routes.SetupServer(server, client)
	// start server
	if err := server.Run(config.Port); err != nil {
		log.Fatalf("failed to start gin server: %v", err.Error())
	}
}
