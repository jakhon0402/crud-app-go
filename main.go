package main

import (
	"crud-app-go/controllers"
	"crud-app-go/initializers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)
	r.Run()
	fmt.Println("Hello")
}
