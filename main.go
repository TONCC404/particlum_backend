package main

import (
	"particlum_backend/config"
	"particlum_backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitDB()
	router.InitRoutes(r)
	r.Run(":8007")
}
