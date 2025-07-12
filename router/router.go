package router

import (
	"particlum_backend/controller"
	"particlum_backend/middleware"

	"github.com/gin-gonic/gin"
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
