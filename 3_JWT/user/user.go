package user

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{

		"message": "GetUser",
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
	})
}

func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{

		"message": "GetUser",
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
	})
}
