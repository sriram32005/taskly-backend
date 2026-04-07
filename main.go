package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sriram32005/taskly-backend/config"
	"github.com/sriram32005/taskly-backend/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Panic("Error during auto migration: ", err)
	}

	r := gin.Default()
	config.ConnectDB()
	r.Run(":"+os.Getenv("PORT"))

}