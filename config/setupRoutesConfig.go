package config

import (
	studentRoute "deneme/student/route"
	teacherRoute "deneme/teacher/route"
	"github.com/gin-gonic/gin"
)

func SetUpAllRoutes(app *gin.Engine) {
	studentRoute.SetupStudentRoutes(app)
	teacherRoute.SetupTeacherRoute(app)
}
