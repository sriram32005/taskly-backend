package handlers

import (
	"net/http"
	"time"

	"github.com/sriram32005/taskly-backend/config"
	"github.com/sriram32005/taskly-backend/models"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Title    string  `json:"title" binding:"required"`
	Priority *string `json:"priority"`  // pointer → detect if provided
	DueDate  *string `json:"due_date"`  // optional
}

func isValidPriority(p string) bool {
	switch p {
	case "low", "medium", "high":
		return true
	default:
		return false
	}
}

func CreateTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	// Default priority
	priority := "medium"

	// If user provided priority
	if input.Priority != nil && *input.Priority != "" {
		if !isValidPriority(*input.Priority) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority"})
			return
		}
		priority = *input.Priority
	}

	// Parse due date (optional)
	var dueDate *time.Time
	if input.DueDate != nil && *input.DueDate != "" {
		t, err := time.Parse(time.RFC3339, *input.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		dueDate = &t
	}

	task := models.Task{
		Title:    input.Title,
		Priority: priority,
		DueDate:  dueDate,
		UserID:   userID,
	}

	config.DB.Create(&task)

	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	userID := c.GetUint("user_id")

	var tasks []models.Task
	config.DB.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var task models.Task

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Toggle complete OR update title
	// var input map[string]interface{}
	var input models.Task
	c.ShouldBindJSON(&input)

	config.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}