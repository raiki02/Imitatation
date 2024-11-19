package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var token = "1234567"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("access_token")
		if accessToken != token {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "token failed",
			})
			c.Abort()
		}
		c.Next()
	}
}
