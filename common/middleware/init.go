package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	//数据库连接
	//r.Use(WithContextDb)
	r.GET("/db", func(c *gin.Context) {
		c.String(200, "db init success")
	})
	//r.Run(":8080")
}
