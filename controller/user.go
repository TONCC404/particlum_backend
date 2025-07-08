package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	// 这里你可以将用户存入数据库（省略 DB 逻辑）
	// 假设成功：
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered",
		"user":    req.Username,
	})
}

// 用户登录接口
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 验证用户名密码逻辑（此处为演示用）
	if req.Username == "admin" && req.Password == "admin" {
		token := "mock-token" // 实际应返回 JWT Token
		c.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// 获取当前用户信息（需要认证的接口）
func GetProfile(c *gin.Context) {
	// 实际应从 JWT Token 中解析出用户信息
	// 此处模拟返回
	user := "admin"
	c.JSON(http.StatusOK, gin.H{
		"username": user,
		"role":     "streamer",
	})
}
