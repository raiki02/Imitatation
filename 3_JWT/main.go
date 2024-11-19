package main

import (
	"jwt/middleware"
	"jwt/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors(), middleware.Cors1())
	routers.InitRouters(r)

	r.Run()
}
