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
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/router"
	"github.com/iyaoo/go-IMChat/common/middleware"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start API server",
	Example: "go-IMChat server",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("read config")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

var AppRouters = make([]func(), 0)

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}
func run() error {
	r := gin.Default()
	sdk.Runtime.SetEngine(r)
	for _, f := range AppRouters {
		f()
	}
	initRouter()

	//启动服务器
	srv := &http.Server{
		Addr:    ":" + "8000",
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
	slog.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	return nil
}

// func initRouter(r *gin.Engine) {
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
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	middleware.InitMiddleware(r)
}
