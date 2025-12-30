package payment_core

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/midsmr/payment/internal/payment_core/application"
)

var rootCmd = &cobra.Command{
	Use:     "payment_core",
	Short:   "Payment is an enterprise-level payment acquiring solution.",
	Version: application.CoreVersion,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 从 flag 中获取 workdir 值
		workdir, _ := cmd.Flags().GetString("workdir")
		if workdir != "" {
			application.Workdir = workdir
		}
	},
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}" + "\nLastCommit: " + application.LastCommit + "\nBuildDate: " + application.BuildDate + "\n")
	// 添加全局 --workdir flag
	rootCmd.PersistentFlags().StringVarP(&application.Workdir, "workdir", "w", ".", "Working directory path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
