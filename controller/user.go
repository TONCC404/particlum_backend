package controller

import (
	"net/http"
	"particlum_backend/auth"

	"github.com/gin-gonic/gin"
)

type PersonalInfo struct {
	role       string   `json:"role"`
	industry   string   `json:"industry"`
	company    string   `json:"company"`
	experience string   `json:"experience"`
	goals      string   `json:"goals"`
	interests  []string `json:"interests"`
	bio        string   `json:"bio"`
}

func Register(c *gin.Context) {
	var req struct {
		email        string       `json:"email" binding:"required"`
		username     string       `json:"username" binding:"required"`
		password     string       `json:"password" binding:"required"`
		PersonalInfo PersonalInfo `json:"personal_info" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := auth.GenerateToken(req.username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// 这里你可以将用户存入数据库（省略 DB 逻辑）
	// 假设成功：
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered",
		"user":    req.username,
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
