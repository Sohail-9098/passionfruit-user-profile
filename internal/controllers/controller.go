package controllers

import (
	"context"
	"net/http"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"github.com/Sohail-9098/passionfruit-user-profile/pkg/database"
	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := database.Client.Database("test").Collection("profiles")
	if _, err := collection.InsertOne(context.TODO(), profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create profile"})
		return
	}
	c.JSON(http.StatusCreated, profile)
}

func GetProfile(c *gin.Context) {
}

func UpdateProfile(c *gin.Context) {
}

func SearchProfiles(c *gin.Context) {
}
