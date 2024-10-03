package main

import (
	"go_crud_api_using_gin/initializers"
	"go_crud_api_using_gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
