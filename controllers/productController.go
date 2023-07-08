package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
)

// product create
func ProductCreate(c *gin.Context) {
	// get data from req body
	var body struct {
		Name       string
		CategoryID int32
		BrandID    int32
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Invalid request body",
		})
		return
	}

	// check if product with the same name already exists
	var existingProduct models.Product
	err = initializers.DB.Where("name = ?", body.Name).First(&existingProduct).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"error":      "Product already exists",
		})
		return
	}

	// create product
	product := models.Product{Name: body.Name, CategoryID: int(body.CategoryID), BrandID: int(body.BrandID)}
	err = initializers.DB.Create(&product).Error

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
		"data":       product,
	})
}

// get all product
func GetAllProduct(c *gin.Context) {
	// get the product
	var product []models.Product
	err := initializers.DB.Find(&product).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "product not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "All product",
		"data":       product,
	})
}

// get single product
func GetSingleProduct(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// get the post
	var product models.Product
	err := initializers.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Product not found",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Product",
		"data":       product,
	})
}

func UpdateProduct(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Name       string
		CategoryID int32
		BrandID    int32
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
	var product models.Product
	err = initializers.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	// update it
	err = initializers.DB.Model(&product).Updates(models.Product{Name: body.Name, CategoryID: int(body.CategoryID), BrandID: int(body.BrandID)}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Update successful",
		"data":       product,
	})
}

// product delete
func DeleteProduct(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")

	// delete it
	var product models.Product
	err := initializers.DB.First(&product, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"error":      err.Error(),
		})
		return
	}

	err = initializers.DB.Delete(&product, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"error":      err.Error(),
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
