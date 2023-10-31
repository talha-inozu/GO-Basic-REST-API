package main

import (
	models "deneme/model"
	"github.com/gin-gonic/gin"
)

func main() {
	models.CreateStudent(15, "Dumy", "Dum", 1)

	app := gin.Default()

	app.GET("/getStudent", func(c *gin.Context) {
		c.JSON(200, models.STUDENT)
	})

	app.POST("/createStudent", func(c *gin.Context) {
		var reqBody models.StudentDTO
		c.ShouldBindJSON(&reqBody)
		models.StudentArray = append(models.StudentArray, reqBody)
		c.JSON(200, models.StudentArray)
	})

	app.Run(":3000")
}
