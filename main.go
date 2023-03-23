package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/gin-json/controllers"
	"github.com/mauFade/gin-json/initializers"
)

// CompileDaemon -command="./gin-json"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)

	r.Run()
}
