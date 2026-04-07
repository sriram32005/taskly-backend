package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sriram32005/taskly-backend/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	r := gin.Default()
	config.ConnectDB()
	r.Run(":"+os.Getenv("PORT"))

}