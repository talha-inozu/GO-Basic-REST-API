package studentRoute

import (
	"deneme/student/model"
	"github.com/gin-gonic/gin"
)

func SetupStudentRoutes(app *gin.Engine) {
	studentService := app.Group("/student")
	{
		studentService.GET("/getStudent", func(c *gin.Context) {
			c.JSON(200, models.STUDENT)
		})

		studentService.POST("/createStudent", func(c *gin.Context) {
			var reqBody models.StudentDTO
			c.ShouldBindJSON(&reqBody)
			models.StudentArray = append(models.StudentArray, reqBody)
			c.JSON(200, models.StudentArray)
		})

	}
}
