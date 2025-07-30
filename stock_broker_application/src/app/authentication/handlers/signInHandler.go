package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var creds models.SignInModel

	// Binding JSON for login credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
		return
	}

	// Validating user from DB
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
	c.JSON(http.StatusOK, gin.H{"message": constants.MsgLoggedIn})
}
