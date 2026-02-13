package main

import (
	"gin-quickstart/db"
	"gin-quickstart/models"
	"gin-quickstart/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDB()

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {

		var users []models.User

		result := db.DB.Find(&users)

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, users)
	})

	routes.SetupRoutes(router)

	router.Run(":8080")
}
