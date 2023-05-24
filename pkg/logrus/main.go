package logrus

import "github.com/siruspen/logrus"

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
