package main

import (
	"crud-app-go/initializers"
	"crud-app-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
