package controllers

import (
	"net/http"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/services"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProfile(c *gin.Context, client *mongo.Client) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		logger.Log(logger.Error, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	id, err := services.AddProfile(client, profile)
	if err != nil {
		logger.Log(logger.Error, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create profile"})
		return
	}
	c.String(http.StatusCreated, id)
}

func UpdateProfile(c *gin.Context, client *mongo.Client) {
	id := c.Params.ByName("id")
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		logger.Log(logger.Error, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if err := services.UpdateProfile(client, profile, id); err != nil {
		logger.Log(logger.Error, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}
