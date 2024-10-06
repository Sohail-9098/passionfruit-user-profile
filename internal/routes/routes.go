package routes

import (
	"github.com/Sohail-9098/passionfruit-user-profile/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/profiles", controllers.CreateProfile)
	r.GET("/profiles/:id", controllers.GetProfile)
	r.PUT("/profiles/:id", controllers.UpdateProfile)
	r.GET("/profiles/search", controllers.SearchProfiles)
}
