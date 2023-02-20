package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	//middleware.InitMiddleware(r)
	// 系统路由
	InitSysRouter(r)
	// 业务路由
	InitProjectsRouter(r)
}
