package payment_core

import (
	"os"

	"github.com/midsmr/payment/internal/payment_core/application"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var startCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP service of payment",
	Run: func(cmd *cobra.Command, args []string) {
		if err := application.InitServer(); err != nil {
			application.Logger.Error("Failed to initialize application", zap.Error(err))
			os.Exit(1)
		}

		defer func() {
			_ = application.Logger.Sync()
		}()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
