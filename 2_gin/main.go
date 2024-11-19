package main

import (
	"fmt"
	"gin_demo/middleware"
	"gin_demo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	//initDB("tcp://127.0.0.1...")
	LoggerMiddleware(initDB)("tcp://127.0.0.1...")
	r.Use(gin.Logger(), gin.Recovery(), middleware.Auth())
	//r.Use(gin.Logger(), gin.Recovery(), CheckAuth2("asdfgh"))
	routers.InitRouters(r)

	r.Run()
}

func CheckAuth(c *gin.Context) {
	fmt.Println("call checkAuth func")
	c.Next()
}

func CheckAuth2(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("call ckeckAuth func", param)
		c.Next()
	}
}

func initDB(connstr string) {
	fmt.Println("init db", connstr)
}

func LoggerMiddleware(in func(connstr string)) func(connstr string) {
	return func(connstr string) {
		fmt.Println("LoggerMiddleware start")
		in(connstr)
		fmt.Println("LoggerMiddleware end")

	}
}
