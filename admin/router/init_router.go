package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/iyaoo/go-IMChat/common/middleware"
	"github.com/iyaoo/reusable-lib/tools/logger"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		logger.Fatal("not support other engine")
		os.Exit(-1)
	}

	r.Use(middleware.RequestId(pkg.TrafficKey))
	middleware.InitMiddleware(r)

	// 业务路由
	InitProjectsRouter(r)
	// 系统路由
	InitSysRouter(r)
}
