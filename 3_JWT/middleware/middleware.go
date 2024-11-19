package middleware

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders: []string{
				"Origin", "Content-Type", "Content-Length",
			},
			AllowMethods: []string{
				"GET", "POST", "PUT", "DELETE",
			},
		})

}

func Cors1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Cors1 is called...")
		c.Next()
	}

}
