package controllers

import (
	"gin-quickstart/db"
	"gin-quickstart/models"
	"gin-quickstart/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateAccRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateAcc(c *gin.Context) {
	var req CreateAccRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if !services.EmailValid(req.Email) {
		c.JSON(400, gin.H{"error": "invalid email"})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	User := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&User).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, gin.H{
		"message":  "success",
		"username": User.Username,
		"email":    User.Email,
	})
}
