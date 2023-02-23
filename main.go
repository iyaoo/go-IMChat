package main

import (
	"github.com/iyaoo/go-IMChat/cmd"
	"github.com/iyaoo/reusable-lib/tools/logger"
)

func main() {
	logger.Info("init logger success")
	cmd.Execute()
}
