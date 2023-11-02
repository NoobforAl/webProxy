package logger

import (
	"errors"
	"io"
	"os"
	"web_proxy/contract"

	"github.com/sirupsen/logrus"
)

var (
	errOpenLogFile = errors.New("open log file error: ")
)

// create new logger
func New(debug bool, pathLogFile string) contract.Logger {
	var (
		logLevel logrus.Level = logrus.InfoLevel
		out      io.Writer    = os.Stderr
		err      error
	)

	// check if path log file is not empty
	// load log file and save log in file
	if pathLogFile != "" {
		flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
		out, err = os.OpenFile(pathLogFile, flag, 0600)
		if err != nil {
			panic(errors.Join(errOpenLogFile, err))
		}
	}

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
