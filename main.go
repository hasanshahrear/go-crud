package main

import (
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
	r.POST("/posts", controllers.PostCreate)

	r.Run()
}
