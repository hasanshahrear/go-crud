package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/controllers"
	"github.com/shahrear/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// public router
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to inventory",
		})
	})

	// brand router
	r.GET("/brands", controllers.GetAllBrand)
	r.GET("/brand/:id", controllers.GetSingleBrand)
	r.POST("/brand/create", controllers.BrandCreate)
	r.PUT("/brand/update/:id", controllers.UpdateBrand)
	r.DELETE("/brand/delete/:id", controllers.DeleteBrand)

	// brand router
	r.GET("/categories", controllers.GetAllCategory)
	r.GET("/category/:id", controllers.GetSingleCategory)
	r.POST("/category/create", controllers.CategoryCreate)
	r.PUT("/category/update/:id", controllers.UpdateCategory)
	r.DELETE("/category/delete/:id", controllers.DeleteCategory)

	// product router
	r.GET("/products", controllers.GetAllProduct)
	r.GET("/product/:id", controllers.GetSingleProduct)
	r.POST("/product/create", controllers.ProductCreate)
	r.PUT("/product/update/:id", controllers.UpdateProduct)
	r.DELETE("/product/delete/:id", controllers.DeleteProduct)

	r.Run()
}
