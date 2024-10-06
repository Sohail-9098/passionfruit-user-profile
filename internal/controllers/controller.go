package controllers

import (
	"context"
	"net/http"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProfile(c *gin.Context, client *mongo.Client) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := client.Database("test").Collection("profiles")
	if _, err := collection.InsertOne(context.TODO(), profile); err != nil {
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
