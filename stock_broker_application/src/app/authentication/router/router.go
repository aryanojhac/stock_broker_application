package router

import (
	"authentication/constants"
	"authentication/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST(constants.SIGNUP, handlers.SignUp)
	router.POST(constants.SIGNIN, handlers.SignIn)
	return router
}
