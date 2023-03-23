package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauFade/gin-json/initializers"
	"github.com/mauFade/gin-json/models"
)

func CreatePost(c *gin.Context) {
	post := models.Post{
		ID:        uuid.NewString(),
		Title:     "example title",
		Body:      "example body",
		CreatedAt: time.Now(),
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)

		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}
