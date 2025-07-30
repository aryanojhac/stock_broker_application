package main

import (
	"authentication/constants"
	"authentication/router"
	"authentication/utils"
	"authentication/utils/db"
)

func main() {
	utils.LoadConfig()
	utils.InitLogger() // initialize loggers

	utils.InfoLogger.Println("Configuration loaded")
	db.ConnectDB()

	utils.InfoLogger.Println("Starting server...")
	r := router.SetupRouter()
	r.Run(constants.PORT)
}
