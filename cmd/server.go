package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/router"
	"github.com/iyaoo/go-IMChat/common/config"
	"github.com/iyaoo/go-IMChat/common/middleware"
	"github.com/iyaoo/reusable-lib/tools/logger"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start API server",
	Example: "go-IMChat server",
	PreRun: func(cmd *cobra.Command, args []string) {
		GetConfigInfos()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

var AppRouters = make([]func(), 0)

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}

// run 启动服务以及路由
func run() error {
	formatter := slog.NewJSONFormatter()
	err := logger.InitLogger(config.App.Config.Settings.Logger.Path, config.App.Config.Settings.Logger.Level, formatter)
	if err != nil {
		logger.Fatalf("init logger err:%v", err)
	}
	switch config.App.Config.Settings.Application.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.TestMode)
	}
	//创建*gin.Engine变量
	r := gin.Default()
	sdk.Runtime.SetEngine(r)

	for _, f := range AppRouters {
		f()
	}
	initRouter()

	StartAndCloseServer(r)
	return nil
}

// initRouter 注册中间件
func initRouter() {
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

	middleware.InitMiddleware(r)
}

// StartAndCloseServer 启、停服务
func StartAndCloseServer(r *gin.Engine) {
	//启动服务器
	srv := &http.Server{
		Addr:    config.App.Config.Settings.Application.Url + ":" + config.App.Config.Settings.Application.Host,
		Handler: sdk.Runtime.GetEngine(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
