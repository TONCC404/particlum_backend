package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"particlum_backend/auth"
)

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// 这里你可以将用户存入数据库（省略 DB 逻辑）
	// 假设成功：
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered",
		"user":    req.Username,
		"token":   token,
	})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if req.Username == "admin" && req.Password == "admin" {
		token, err := auth.GenerateToken(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func GetProfile(c *gin.Context) {
	// 实际应从 JWT Token 中解析出用户信息
	// 此处模拟返回
	user := "admin"
	c.JSON(http.StatusOK, gin.H{
		"username": user,
		"role":     "streamer",
	})
}
