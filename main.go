package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sriram32005/taskly-backend/config"
	"github.com/sriram32005/taskly-backend/models"
	"github.com/sriram32005/taskly-backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	
	r := gin.Default()
	config.ConnectDB()
	
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Panic("Error during auto migration: ", err)
	}

	// Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)


	r.Run(":"+os.Getenv("PORT"))

}