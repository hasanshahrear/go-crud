package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

// category create
func CategoryCreate(c *gin.Context) {
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

	// check if category with the same name already exists
	var existingCategory models.Category
	err = initializers.DB.Where("name = ?", body.Name).First(&existingCategory).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"error":      "Category already exists",
		})
		return
	}

	// create category
	category := models.Category{Name: body.Name}
	err = initializers.DB.Create(&category).Error

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
		"data":       category,
	})
}

// get all category
func GetAllCategory(c *gin.Context) {
	// get the category
	var category []models.Category
	err := initializers.DB.Find(&category).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "category not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "All category",
		"data":       category,
	})
}

// get single category
func GetSingleCategory(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// get the post
	var category models.Category
	err := initializers.DB.First(&category, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Category not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Category",
		"data":       category,
	})
}

func UpdateCategory(c *gin.Context) {
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
	var category models.Category
	err = initializers.DB.First(&category, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Category not found",
		})
		return
	}

	// update it
	err = initializers.DB.Model(&category).Updates(models.Category{Name: body.Name}).Error
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
		"data":       category,
	})
}

// category delete
func DeleteCategory(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")

	// delete it
	var category models.Category
	err := initializers.DB.First(&category, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"error":      "Category not found",
		})
		return
	}

	err = initializers.DB.Delete(&category, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"error":      "Category not deleted",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Delete successful",
		"data":       nil,
	})
}
