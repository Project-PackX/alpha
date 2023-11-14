package utils

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})

	return logger
}

var Logger = SetupLogger()
