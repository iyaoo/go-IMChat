package middleware

import (
	"github.com/gin-gonic/gin"
)

// InitMiddleware 引用中间件
func InitMiddleware(r *gin.Engine) {
	//数据库连接
	//r.Use(WithContextDb)
	r.Use(LoggerToFile())
}
