package controllers

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginAcc(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "username and password are required"})
		return
	}

	user := models.User{}
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	c.JSON(200, gin.H{
		"message":  "success",
		"username": user.Username,
		"email":    user.Email,
	})
}
