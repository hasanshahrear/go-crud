package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

func BrandCreate(c *gin.Context) {
	// get data from req body
	var body struct {
		Name string
	}
	c.Bind(&body)

	// create brand
	brand := models.Brand{Name: body.Name}
	result := initializers.DB.Create(&brand)

	// return response
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"error":      result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"data":       brand,
	})
}
