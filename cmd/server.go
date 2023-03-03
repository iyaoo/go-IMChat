package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/router"
	"github.com/iyaoo/go-IMChat/common/config"
	"github.com/iyaoo/reusable-lib/tools/logger"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start API server",
	Example: "go-IMChat server",
	PreRun: func(cmd *cobra.Command, args []string) {
		InitConfigAndLogger()
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
	engine := sdk.Runtime.GetEngine()
	if engine == nil {
		engine = gin.New()
	}

	switch config.App.Config.Settings.Application.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.TestMode)
	}

	err := logger.InitLogger(config.App.Config.Settings.Logger.Path, config.App.Config.Settings.Logger.Level, map[string]interface{}{
		"appname":  config.App.Config.Settings.Application.Name,
		"hostname": pkg.GetLocaHonst(),
	}, slog.NewJSONFormatter())
	if err != nil {
		logger.Fatalf("init logger err %v", err)
	}

	for _, f := range AppRouters {
		f()
	}

	StartAndCloseServer()

	return nil
}

// StartAndCloseServer 启、停服务
func StartAndCloseServer() {
	//启动服务器
	srv := &http.Server{
		Addr:    config.App.Config.Settings.Application.Url + ":" + config.App.Config.Settings.Application.Port,
		Handler: sdk.Runtime.GetEngine(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Fatalf("listen: %s\n", err)
		}
	}()
	logger.Info("Service started successfully!")
	fmt.Println(pkg.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.App.Config.Settings.Application.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", pkg.GetLocaHonst(), config.App.Config.Settings.Application.Port)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown Server ... \r\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
