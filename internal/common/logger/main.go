package logger

import "github.com/siruspen/logrus"

type LoggerInterface interface {
	Error(err error, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
}

type logger struct {
	log *logrus.Logger
}

func New(logrus *logrus.Logger) LoggerInterface {
	return &logger{log: logrus}
}

func (l *logger) Error(err error, args ...interface{}) {
	l.log.Errorf("%s\tcontent: %+v", err.Error(), args)
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.log.Infof("%s\tcontent: %+v", msg, args)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.log.Warnf("%s\tcontent: %+v", msg, args)
}
