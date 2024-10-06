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
	database.Connect(MONGODB_URI)
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(PORT)
}
