package main

import (
	"github.com/gin-gonic/gin"
	"particlumn_backend/router"
)

func main() {
	r := gin.Default()
	router.InitRoutes(r)
	r.Run(":8007")
}
