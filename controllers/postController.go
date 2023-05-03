package controllers

import (
	"crud-app-go/initializers"
	"crud-app-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	result := initializers.DB.Create(&models.Post{Id: uuid.New(), Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}
