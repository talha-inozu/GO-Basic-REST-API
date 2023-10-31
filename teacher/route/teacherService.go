package teacherRoute

import (
	"deneme/teacher/model"
	"github.com/gin-gonic/gin"
)

func SetupTeacherRoute(app *gin.Engine) {
	studentService := app.Group("/teacher")
	{
		studentService.GET("/getTeacher", func(c *gin.Context) {
			c.JSON(200, models.Teacher)
		})

		studentService.POST("/createTeacher", func(c *gin.Context) {
			var reqBody models.TeacherDTO
			c.ShouldBindJSON(&reqBody)
			models.TeacherArray = append(models.TeacherArray, reqBody)
			c.JSON(200, models.TeacherArray)
		})

	}
}
