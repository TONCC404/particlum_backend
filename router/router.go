package router

import (
	"particlum_backend/controller"
	"particlum_backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/profile", controller.GetProfile)
	}
}
