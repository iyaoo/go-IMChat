package router

import (
	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(v1 *gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

// 初始化业务路由
func InitProjectsRouter(r *gin.Engine) *gin.Engine {

	// 需要认证的路由
	examplesCheckRoleRouter(r)

	// 无需认证的路由
	examplesNoCheckRoleRouter(r)
	return r
}

// examplesNoCheckRoleRouter 启动所有无需认证路由
func examplesNoCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group("/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// examplesCheckRoleRouter 启动所有需要认证路由
func examplesCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v1)
	}
}
