package router

import (
	"github.com/gin-gonic/gin"
)

// InitSysRouter 初始化系统路由
func InitSysRouter(r *gin.Engine) {
	r.GET("/sys", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "system router initialization",
		})
	})
}
