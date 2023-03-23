package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/gin-json/initializers"
	"github.com/mauFade/gin-json/models"
)

func CreatePost(c *gin.Context) {
	post := models.Post{
		Title: "example title",
		Body:  "example body",
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
