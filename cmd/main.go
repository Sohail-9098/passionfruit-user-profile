package main

import (
	"github.com/Sohail-9098/passionfruit-user-profile/internal/routes"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/database"
	"github.com/gin-gonic/gin"
)

const (
	MONGODB_URI = "mongodb://localhost:27017"
	PORT        = ":8080"
)

func main() {
	client := database.Connect(MONGODB_URI)
	server := gin.Default()
	routes.SetupServer(server, client)
	server.Run(PORT)
}
