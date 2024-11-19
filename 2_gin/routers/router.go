package routers

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	InitCourse(r)
	InitUser(r)
}
