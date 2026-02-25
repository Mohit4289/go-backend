package todo

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddTodoRequest struct {
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

func AddTodo(c *gin.Context) {
	var req AddTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id format"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	var existingTodo models.Todo
	if err := db.DB.Where("text = ? AND user_id = ?", req.Text, userID).First(&existingTodo).Error; err == nil {
		c.JSON(400, gin.H{"error": "todo already exists"})
		return
	}

	todo := models.Todo{
		Text:   req.Text,
		UserID: userID,
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo created successfully"})
}
