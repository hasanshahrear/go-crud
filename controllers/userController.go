package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahrear/go-crud/initializers"
	"github.com/shahrear/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email and password from body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "SignUp Successful",
	})
}

func Login(c *gin.Context) {
	// get the email & password from the req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	// look up the requested user

	// compare the send password

	// generate jwt token

	// response
}
