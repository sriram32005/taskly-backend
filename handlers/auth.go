package handlers

import (
	"net/http"

	"github.com/sriram32005/taskly-backend/config"
	"github.com/sriram32005/taskly-backend/models"
	"github.com/sriram32005/taskly-backend/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := utils.HashPassword(input.Password)

	user := models.User{
		Email:        input.Email,
		PasswordHash: hash,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Login(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := utils.CheckPassword(input.Password, user.PasswordHash); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func APIInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Taskly API usage",
		"usage": []gin.H{
			{
				"endpoint":    "POST /register",
				"description": "Register a new user",
				"body":        gin.H{"email": "user@example.com", "password": "secret"},
			},
			{
				"endpoint":    "POST /login",
				"description": "Authenticate and receive a Bearer token",
				"body":        gin.H{"email": "user@example.com", "password": "secret"},
			},
			{
				"endpoint":    "GET /endpoint",
				"description": "Get this API usage information",
			},
			{
				"endpoint":    "POST /tasks",
				"description": "Create a new task (authenticated)",
				"headers":     gin.H{"Authorization": "Bearer <token>"},
				"body":        gin.H{"title": "Buy groceries", "priority": "high", "due_date": "2026-05-01T15:00:00Z"},
			},
			{
				"endpoint":    "GET /tasks",
				"description": "List tasks for the authenticated user",
				"headers":     gin.H{"Authorization": "Bearer <token>"},
			},
			{
				"endpoint":    "PUT /tasks/:id",
				"description": "Update a task by ID (authenticated)",
				"headers":     gin.H{"Authorization": "Bearer <token>"},
				"body":        gin.H{"title": "Updated title", "priority": "medium"},
			},
			{
				"endpoint":    "DELETE /tasks/:id",
				"description": "Delete a task by ID (authenticated)",
				"headers":     gin.H{"Authorization": "Bearer <token>"},
			},
		},
	})
}
