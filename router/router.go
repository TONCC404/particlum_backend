package router

import (
	"github.com/gin-gonic/gin"
	"particlumn_backend/controller"
	"particlumn_backend/middleware"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware()) // 这里加保护
	{
		auth.GET("/profile", controller.GetProfile)
	}
}
