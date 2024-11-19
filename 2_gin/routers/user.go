package routers

import (
	"gin_demo/web"

	"github.com/gin-gonic/gin"
)

func InitUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/user", web.GetUser)
	v1.POST("/user", web.AddUser)
	v1.PUT("/user")

}
