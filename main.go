package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sriram32005/taskly-backend/config"
	"github.com/sriram32005/taskly-backend/models"
	"github.com/sriram32005/taskly-backend/handlers"
	"github.com/sriram32005/taskly-backend/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	
	r := gin.Default()
	config.ConnectDB()
	
	if err := config.DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Panic("Error during auto migration: ", err)
	}

	// Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/tasks", handlers.CreateTask)
		auth.GET("/tasks", handlers.GetTasks)
		auth.PUT("/tasks/:id", handlers.UpdateTask)
		auth.DELETE("/tasks/:id", handlers.DeleteTask)
	}


	r.Run(":"+os.Getenv("PORT"))

}