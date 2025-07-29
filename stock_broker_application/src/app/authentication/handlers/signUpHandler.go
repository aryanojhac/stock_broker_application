package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"authentication/models"
	"authentication/service"
	"authentication/constants"              // using for errors
)

// SignUp handles multiple users in a single request
func SignUp(c *gin.Context) {
	// var single models.User
	var multiple []models.User

	if err := c.ShouldBindJSON(&multiple); err == nil {
		for _, user := range multiple {
			if err := service.RegisterUser(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),    //Using actual error from service
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "all users registered successfully"})
		return
	}
	// if err := c.ShouldBindJSON(&single); err == nil {
	// 	if err := service.RegisterUser(single); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
	// 	return
	// }

	c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
}


// SignIn handles login for one user
func SignIn(c *gin.Context) {
	var creds models.User

	// 🔍 Bind JSON for login credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
		return
	}

	// Validate user from DB
	valid, err := service.LoginUser(creds.UserName, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUserNotFound})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidPassword})
		return
	}

	// Successful login
	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully"})
}
