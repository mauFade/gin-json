package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauFade/gin-json/initializers"
	"github.com/mauFade/gin-json/models"
)

type CreatepostInputDTO struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePost(c *gin.Context) {
	var body CreatepostInputDTO

	c.Bind(&body)

	post := models.Post{
		ID:        uuid.NewString(),
		Title:     body.Title,
		Body:      body.Body,
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
