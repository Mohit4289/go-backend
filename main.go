package main

import (
	"gin-quickstart/controllers"
	"gin-quickstart/routes"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/User", func(c *gin.Context) {

		if len(controllers.AccStore) == 0 {
			c.JSON(404, gin.H{
				"error": "user not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"user":    controllers.AccStore,
		})
	})

	routes.SetupRoutes(router)

	router.Run(":8080")
}
