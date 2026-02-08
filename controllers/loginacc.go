package controllers

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginAcc(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "email and password are required",
		})
		return
	}

	// Check if email exists in AccStore
	var foundUser *CreateAccRequest
	for _, user := range AccStore {
		if user.Email == req.Email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	if foundUser.Password != req.Password {
		c.JSON(401, gin.H{
			"error": "invalid password",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "login successful",
		"username": foundUser.Username,
		"email":    foundUser.Email,
	})
}
