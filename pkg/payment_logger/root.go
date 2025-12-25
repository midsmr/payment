package payment_logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(path string) (*zap.Logger, error) {

	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(path, "app.log")

	consoleCore := newConsoleCore()
	fileCore, err := newFileCore(filePath)

	if err != nil {
		return nil, err
	}

	return zap.New(zapcore.NewTee(consoleCore, fileCore), zap.AddCaller()), nil
}
