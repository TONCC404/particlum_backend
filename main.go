package main

import (
	"github.com/gin-gonic/gin"
	"particlumn_backend/config"
	"particlumn_backend/router"
)

func main() {
	r := gin.Default()
	config.InitDB()
	router.InitRoutes(r)
	r.Run(":8007")
}
