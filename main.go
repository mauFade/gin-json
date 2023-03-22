package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/gin-json/initializers"
)

// CompileDaemon -command="./gin-json"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
