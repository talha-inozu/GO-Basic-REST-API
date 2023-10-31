package main

import (
	"deneme/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.SetUpAllRoutes(app)
	app.Run(":8080")
}
