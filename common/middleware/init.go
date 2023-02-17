package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	//数据库连接
	r.Use(WithContextDb)
}
