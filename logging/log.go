package logging

import (
	"github.com/sirupsen/logrus"
)

const (
	// Sys defines the field denoting the system when creating a new logger
	Sys = "sys"

	// Subsys defines the field denoting the subsystem when creating a new logger
	Subsys = "subsys"
)

// Logger is the default base logger
var Logger = defaultLogger()

func defaultLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableTimestamp: true,
		DisableColors:    true,
		FullTimestamp:    true,
	}
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

// SetLogLevel sets the log level on Logger
func SetLogLevel(level logrus.Level) {
	Logger.SetLevel(level)
}
