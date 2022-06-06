package logger

import (
	"context"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func Default() *logrus.Logger {
	defaultLog := logrus.New()
	defaultLog.SetFormatter(&logrus.TextFormatter{})
	level, err := logrus.ParseLevel("info")
	if err != nil {
		// default level to info
		level = logrus.InfoLevel
	}
	defaultLog.SetLevel(level)
	return defaultLog
}

func NoopLogger() *logrus.Logger {
	noopLogger := logrus.New()
	noopLogger.SetOutput(ioutil.Discard)
	return noopLogger
}

func Init(logger *logrus.Logger) {
	log = logger
}

func WithContext(ctx context.Context) *logrus.Entry {
	return log.WithContext(ctx)
}

func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}
