package application

import (
	"os"
	"path/filepath"

	"github.com/midsmr/payment/pkg/payment_logger"
)

func Init() error {

	if err := initWorkdir(); err != nil {
		return err
	}

	if err := initLogger(); err != nil {
		return err
	}

	Logger.Info("Application running in root path: " + Workdir)

	Logger.Info("Application initialized successfully.")

	return nil
}

func initWorkdir() error {
	// 使用 Workdir 的值（默认为 "." 或从 cobra 命令行接收）
	var absPath string
	var err error

	// 如果 Workdir 为 "."，则使用当前工作目录
	if Workdir == "." {
		absPath, err = os.Getwd()
	} else {
		// 否则将提供的路径转换为绝对路径
		absPath, err = filepath.Abs(Workdir)
	}

	if err != nil {
		return err
	}

	Workdir = absPath
	return nil
}

func initLogger() error {
	path := filepath.Join(Workdir, "logs")

	log, err := payment_logger.Init(path)

	if err != nil {
		return err
	}

	Logger = log

	Logger.Info("Initializing application...")

	Logger.Info("Logger initialized successfully. File path： " + path)

	return nil
}
