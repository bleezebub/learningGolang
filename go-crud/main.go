package main

import (
	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/", controllers.PostCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
