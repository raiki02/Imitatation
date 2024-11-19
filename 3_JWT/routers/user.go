package routers

import (
	"jwt/user"

	"github.com/gin-gonic/gin"
)

func initUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/user", user.GetUser)
		v1.POST("/user", user.AddUser)
	}
	v2 := group.Group("/v2")
	{
		v2.GET("/user", user.GetUserV2)
		v2.POST("/user", user.AddUserV2)

	}
}
