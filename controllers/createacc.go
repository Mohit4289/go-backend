package controllers

import (
	"gin-quickstart/services"

	"github.com/gin-gonic/gin"
)

type CreateAccRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var AccStore = make(map[string]CreateAccRequest)

func CreateAcc(c *gin.Context) {
	var req CreateAccRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json",
		})
		return
	}

	var validEmail = services.EmailValid(req.Email)

	if !validEmail {
		c.JSON(400, gin.H{
			"error": "invalid email",
		})
		return
	}

	if _, exists := AccStore[req.Username]; exists {
		c.JSON(400, gin.H{
			"error": "username already exists",
		})
		return
	}

	AccStore[req.Username] = req

	c.JSON(200, gin.H{
		"message":  "success",
		"username": req.Username,
		"email":    req.Email,
		"password": req.Password,
	})

}
