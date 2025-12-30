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
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}" + "\nLastCommit: " + application.LastCommit + "\nBuildDate: " + application.BuildDate + "\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
