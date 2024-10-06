package main

import (
	"context"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/config"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/routes"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/database"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup logger
	logger.SetupLogger()
	// connect to database
	logger.Log(logger.Info, "database setup initiated")
	client := database.Connect(config.MongoDbUri)
	defer client.Disconnect(context.TODO())
	logger.Log(logger.Info, "database setup completed")
	// setup server
	logger.Log(logger.Info, "server setup initiated")
	server := gin.Default()
	routes.SetupServer(server, client)
	logger.Log(logger.Info, "database setup completed")
	// start server
	if err := server.Run(config.Port); err != nil {
		logger.Log(logger.Error, "failed to start gin server: "+err.Error())
	}
}
