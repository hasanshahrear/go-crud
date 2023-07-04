package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

// brand create
func BrandCreate(c *gin.Context) {
	// get data from req body
	var body struct {
		Name string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Invalid request body",
		})
		return
	}

	// create brand
	brand := models.Brand{Name: body.Name}
	err = initializers.DB.Create(&brand).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusInternalServerError,
			"error":      "Create failed",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       brand,
	})
}

// get all brands
func GetAllBrand(c *gin.Context) {
	// get the brand
	var brands []models.Brand
	err := initializers.DB.Find(&brands).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Brands not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "All Brands",
		"data":       brands,
	})
}

// get single brand
func GetSingleBrand(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// get the post
	var brand models.Brand
	err := initializers.DB.First(&brand, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Brand not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Brand",
		"data":       brand,
	})
}

func UpdateBrand(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Name string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Invalid request body",
		})
		return
	}

	// find where to update it
	var brand models.Brand
	err = initializers.DB.First(&brand, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Brand not found",
		})
		return
	}

	// update it
	err = initializers.DB.Model(&brand).Updates(models.Brand{Name: body.Name}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Update failed",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Update successful",
		"data":       brand,
	})
}

// brand delete
func DeleteBrand(c *gin.Context) {
	// get the id from  url
	id := c.Param("id")

	// delete it
	var brand models.Brand
	err := initializers.DB.Delete(&brand, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Brand not found",
		})
		return
	}

	// return resonse
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Delete successful",
		"data":       nil,
	})
}
