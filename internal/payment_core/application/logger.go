package application

import (
	"os"
	"path/filepath"

	"github.com/midsmr/payment/pkg/payment_logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = payment_logger.NewConsoleLogger()

func initLogger() error {
	path := filepath.Join(Workdir, "logs")

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(path, "app.log")

	consoleCore := payment_logger.NewConsoleCore()
	fileCore, err := payment_logger.NewFileCore(filePath)

	if err != nil {
		return err
	}

	Logger = zap.New(zapcore.NewTee(consoleCore, fileCore), zap.AddCaller())

	return nil
}
