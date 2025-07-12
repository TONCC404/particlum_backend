package controller

import (
	"net/http"
	"particlum_backend/auth"
	"particlum_backend/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
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
		Email         string             `json:"email" binding:"required"`
		Username      string             `json:"username" binding:"required"`
		Password      string             `json:"password" binding:"required"`
		Personal_data model.PersonalInfo `json:"personal_data" binding:"required"`
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

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	passwordHash := string(passwordHashBytes)

	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
		PersonalInfo: req.Personal_data,
	}

	if err := model.CreateUser(&user); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "users_username_key":
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			case "users_email_key":
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "Duplicate value"})
			}
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered",
		"user":    req.Username,
		"token":   token,
	})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if req.Email != "" && req.Password != "" {
		token, err := auth.GenerateToken(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
			return
		}
		user, err := model.FindUserByEmail(req.Email)
		if err != nil || !auth.CheckPassword(user.PasswordHash, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"user":    user,
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
