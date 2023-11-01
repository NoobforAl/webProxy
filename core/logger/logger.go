package logger

import (
	"io"
	"os"
	"web_proxy/core/contract"

	"github.com/sirupsen/logrus"
)

func New(debug bool, logFile string) contract.Logger {
	var out io.Writer = os.Stderr
	var err error

	if logFile != "" {
		out, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
	}

	logLevel := logrus.InfoLevel
	if debug {
		logLevel = logrus.DebugLevel
	}

	return &logrus.Logger{
		Out:       out,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logLevel,
	}
}
