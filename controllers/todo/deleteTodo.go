package todo

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "please provide an id"})
		return
	}

	if err := db.DB.Where("id = ?", id).Delete(&models.Todo{}).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to delete todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo deleted successfully"})
}
