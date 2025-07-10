package main

import (
	"github.com/feribeirods/voz-api/config"
	"github.com/feribeirods/voz-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
