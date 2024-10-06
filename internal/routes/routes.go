package routes

import (
	"github.com/Sohail-9098/passionfruit-user-profile/internal/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, client *mongo.Client) {
	r.POST("/profiles", handleRequest(controllers.CreateProfile, client))
	r.GET("/profiles/:id", handleRequest(controllers.GetProfile, client))
	r.PUT("/profiles/:id", handleRequest(controllers.UpdateProfile, client))
	r.GET("/profiles/search", handleRequest(controllers.SearchProfiles, client))
}

func handleRequest(handler func(c *gin.Context, i *mongo.Client), client *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c, client)
	}
}
