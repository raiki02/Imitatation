package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	id := context.DefaultQuery("id", "0")
	//id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{"id": id})
}

type User struct {
	Name  string `form:"name" binding": "required"`
	Phone string `form:"phone" binding": "required , e164"`
}

func AddUser(context *gin.Context) {
	req := &User{}
	err := context.ShouldBind(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, req)
}
