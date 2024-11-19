package user

import "github.com/gin-gonic/gin"

func GetUserV2(c *gin.Context) {
	c.JSON(200, gin.H{

		"message": "GetUser",
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
	})
}

func AddUserV2(c *gin.Context) {
	c.JSON(200, gin.H{

		"message": "GetUser",
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
	})
}
