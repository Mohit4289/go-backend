package controllers

import (
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type CreateAccRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func CreateAcc(c *gin.Context) {
	var req CreateAccRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json or format is wrong"})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "username already exists"})
		return
	}

	var existingEmail models.User
	if exist := db.DB.Where("email = ?", req.Email).First(&existingEmail).Error; exist == nil {
		c.JSON(400, gin.H{"error": "email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	userID := uuid.New()

	User := models.User{
		ID:       userID,
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
