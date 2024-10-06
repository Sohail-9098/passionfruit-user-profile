package controllers

import (
	"log"
	"net/http"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProfile(c *gin.Context, client *mongo.Client) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if err := services.AddProfile(profile, client); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create profile"})
		return
	}
	c.JSON(http.StatusCreated, profile)
}

func GetProfile(c *gin.Context, client *mongo.Client) {
}

func UpdateProfile(c *gin.Context, client *mongo.Client) {
}

func SearchProfiles(c *gin.Context, client *mongo.Client) {
}
