package application

import (
	"github.com/midsmr/payment/pkg/payment_logger"
)

var Workdir string = "."

var Logger = payment_logger.NewConsoleLogger()

var CoreVersion = "dev"

var LastCommit = "unknown"

var BuildDate = "unknown"
