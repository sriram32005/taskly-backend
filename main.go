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
	"github.com/gin-contrib/cors"

)

func main() {
	if os.Getenv("ENV") != "production" {
		godotenv.Load(".env")
	}
	
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://172.17.0.4:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
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