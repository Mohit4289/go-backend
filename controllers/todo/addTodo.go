package todo

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if err := db.DB.First(&models.User{}, todo.UserID).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	if err := db.DB.First(&models.Todo{}, "text = ?", todo.Text).Error; err == nil {
		c.JSON(400, gin.H{"error": "todo already exists"})
		return
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo created successfully"})
}
