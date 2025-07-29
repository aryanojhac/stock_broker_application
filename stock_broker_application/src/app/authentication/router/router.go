package router

import (
	"github.com/gin-gonic/gin"
	"authentication/handlers"
	"authentication/constants"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST(constants.SIGNUP, handlers.SignUp)
	router.POST(constants.SIGNIN, handlers.SignIn)
	return router
}