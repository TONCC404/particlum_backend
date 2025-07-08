package router

import (
	"github.com/gin-gonic/gin"
	"particlumn_backend/controller"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})
	user := r.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/profile", controller.GetProfile)
	}
}
