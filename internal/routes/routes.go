package routes

import (
	"github.com/Sohail-9098/passionfruit-user-profile/internal/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupServer(r *gin.Engine, db *mongo.Client) {
	r.POST("/profiles", handleRequest(controllers.CreateProfile, db))
	r.GET("/profiles/:id", handleRequest(controllers.GetProfile, db))
	r.PUT("/profiles/:id", handleRequest(controllers.UpdateProfile, db))
	r.GET("/profiles/search", handleRequest(controllers.SearchProfiles, db))
}

func handleRequest(handler func(c *gin.Context, i *mongo.Client), client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, client)
	}
}
