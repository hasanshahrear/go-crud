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
	initializers.DB.AutoMigrate(&models.Brand{})
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Purchase{})
	initializers.DB.AutoMigrate(&models.Stock{})
	initializers.DB.AutoMigrate(&models.Sales{})
}
