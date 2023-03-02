package cmd

import (
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/common/config"
	"github.com/iyaoo/reusable-lib/sdk/pkg"
	"github.com/iyaoo/reusable-lib/tools/logger"
)

func InitConfigAndLogger() {
	//初始化配置文件
	config.InitConfig()
	//初始化logger
	formatter := slog.NewJSONFormatter()
	err := logger.InitLogger(config.App.Config.Settings.Logger.Path, config.App.Config.Settings.Logger.Level, map[string]interface{}{
		"appname":  config.App.Config.Settings.Application.Name,
		"hostname": pkg.GetLocalHost(),
		"method":   "",
		"uri":      "",
	}, formatter)
	if err != nil {
		logger.Fatalf("init logger err:%v", err)
	}
	logger.Info("InitLogger success")
}
