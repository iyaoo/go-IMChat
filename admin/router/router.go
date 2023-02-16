package router

import (
	"github.com/gin-gonic/gin"
)

var (
	routerCheckRole = make([]func(v1 *gin.RouterGroup), 0)
)

func InitExamplesRouter(r *gin.Engine) *gin.Engine {

	// 无需认证的路由
	CheckRoleRouter(r)
	return r
}
func CheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v1)
	}
}
