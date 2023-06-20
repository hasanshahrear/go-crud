package main

import (
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
