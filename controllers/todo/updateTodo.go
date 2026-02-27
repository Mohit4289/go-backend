package todo

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
)

type UpdateTodoRequest struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

func UpdateTodo(c *gin.Context) {
	var req UpdateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if req.Status != "pending" && req.Status != "completed" {
		c.JSON(400, gin.H{"error": "invalid status"})
		return
	}

	var existingTodo models.Todo
	if err := db.DB.First(&existingTodo, req.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	existingTodo.Text = req.Text
	existingTodo.Status = req.Status

	if err := db.DB.Save(&existingTodo).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo updated successfully"})
}
