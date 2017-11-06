package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func DefaultLogger() *logrus.Logger {
	return log
}
