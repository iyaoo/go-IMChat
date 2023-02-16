package cmd

import (
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/common/gorm"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Use cobra to connect to the database",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := gorm.InitGorm()
		if err != nil {
			slog.Error("gorm connect mysql err:%s", err)
		}
	},
}
