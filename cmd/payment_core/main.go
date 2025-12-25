package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/midsmr/payment/internal/payment_core/application"
)

var rootCmd = &cobra.Command{
	Use:     "payment_core",
	Short:   "start payment core service",
	Version: application.CoreVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := application.Init(); err != nil {
			return err
		}

		defer func() {
			// Sync logger but don't fail on sync errors (common on Windows with stdout)
			_ = application.Logger.Sync()
		}()

		return nil
	},
}

func main() {
	rootCmd.SetVersionTemplate("Version: {{.Version}}" + "\nLastCommit: " + application.LastCommit + "\nBuildDate: " + application.BuildDate + "\n")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
