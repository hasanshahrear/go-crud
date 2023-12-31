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

	// user routers
	r.POST("/sign-up", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// todo routers
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
