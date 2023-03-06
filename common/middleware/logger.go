package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iyaoo/reusable-lib/tools/logger"
)

// LoggerToFile gin日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求

		c.Next()
		// 结束时间
		endTime := time.Now()

		reqUri := c.Request.RequestURI
		logData := map[string]any{
			"statusCode":  c.Writer.Status(),
			"latencyTime": endTime.Sub(startTime),
			"clientIP":    c.ClientIP(),
			"method":      c.Request.Method,
			"uri":         c.Request.RequestURI,
		}
		if reqUri != "/info" {
			logger.Info(logData)
		}
	}
}
