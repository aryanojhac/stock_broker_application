package main

import (
	"authentication/utils"
	"authentication/utils/db"
	"authentication/router"
)

func main() {
	utils.LoadConfig()
	utils.InitLogger() // initialize loggers

	utils.InfoLogger.Println("Configuration loaded")
	db.ConnectDB()

	utils.InfoLogger.Println("Starting server...")
	r := router.SetupRouter()
	r.Run(":" + utils.AppConfig.ServerPort)
}
