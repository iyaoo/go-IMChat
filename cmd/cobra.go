package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-IMChat",
	Short: "go-IMChat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(startCmd)
}
