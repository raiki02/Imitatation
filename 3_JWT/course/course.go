package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	Name     string `json:"name" form:"name"`
	Teacher  string `json:"teacher" form:"teacher"`
	Duration string `json:"duration" form:"duration" ` //binding:"required , number"????
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Get Course",
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
	})
}

func PostCourse(c *gin.Context) {
	req := &course{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}

func PutCourse(c *gin.Context) {
	req := &course{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}

func DeleteCourse(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
