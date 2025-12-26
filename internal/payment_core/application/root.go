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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	Workdir = cwd
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
