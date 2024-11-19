package routers

import (
	"jwt/course"

	"github.com/gin-gonic/gin"
)

func initCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/course/:id", course.GetCourse)

		v1.POST("/course", course.PostCourse)

		v1.PUT("/course", course.PutCourse)

		v1.DELETE("course", course.DeleteCourse)
	}
}
