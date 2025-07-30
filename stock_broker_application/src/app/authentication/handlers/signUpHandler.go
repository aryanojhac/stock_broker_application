package handlers

import (
	"authentication/constants" // using for errors
	"authentication/models"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp handles multiple users in a single request
func SignUp(c *gin.Context) {
	// var single models.User
	var multiple []models.SignUpModel

	if err := c.ShouldBindJSON(&multiple); err == nil {
		for _, user := range multiple {
			if err := service.RegisterUser(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(), //Using actual error from service
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": constants.MsgRegistered})
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
