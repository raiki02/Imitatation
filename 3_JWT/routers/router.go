package routers

import (
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")

	initCourse(api)
	initUser(api)

	notAuthApi := r.Group("/api")
	notAuthApi.Use(middleware.Cors())
	initLogin(api)
}
