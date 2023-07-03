package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

func PostCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the post
	var posts []models.Post

	initializers.DB.Find(&posts)

	//response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get th post
	var post models.Post
	initializers.DB.First(&post, id)

	// response with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get the data  off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// find the post where update it
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	var post models.Post

	//find the post
	initializers.DB.Delete(&post, id)

	//response with it
	c.Status(200)

}
