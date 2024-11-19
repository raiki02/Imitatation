package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(context *gin.Context) {
	id := context.DefaultQuery("id", "0")
	//id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{"id": id})
}

type Course struct {
	Name    string `form:"name" binding": "required"`
	Teacher string `form:"teacher" binding": "required"`
	Price   string `form:"price" binding": "number"`
}

func AddCourse(context *gin.Context) {
	req := &Course{}
	err := context.ShouldBind(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, req)
}
