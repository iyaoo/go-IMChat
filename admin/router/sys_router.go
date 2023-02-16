package router

import (
	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine) {
	r.GET("/sys", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "system router initialization",
		})
	})
	//r.Run(":8080")
}
