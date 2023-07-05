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

	r.Run()
}
