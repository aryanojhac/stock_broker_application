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
	var multiple []models.SignUpModel

	// Binding the JSON input
	if err := c.ShouldBindJSON(&multiple); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrInvalidRequest.Error(), // Proper error if JSON is malformed
		})
		return
	}

	// Manually validating for required fields
	for _, user := range multiple {
		
		if user.UserName == "" || user.Password == "" || user.Email == "" || user.PhoneNumber == "" || user.PanCard == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "all fields are required"})
			return
		}

		// Calling service layer to register the user
		if err := service.RegisterUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(), // Return the actual error from service/repo
			})
			return
		}
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": constants.MsgRegistered})
}
