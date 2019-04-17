package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logFile *os.File
var logLevel log.Level

//MakeLogger logger generate
func MakeLogger() *log.Entry {
	var logger *log.Entry
	return logger
}

//NewLogConf for gloal log config
func NewLogConf(logPath string, level string) {
	var err error
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0750)
	if err != nil {
		log.WithField("logFrom", "utils.log").Error(err)
		os.Exit(1)
	}

	switch level {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "warn":
		logLevel = log.WarnLevel
	case "error":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.WarnLevel
	}
}

//Logger for devops app
func Logger(name string) *log.Entry {
	var out = logFile
	if logLevel == log.DebugLevel {
		out = os.Stdout
	}

	logger := log.Logger{
		Formatter: &log.TextFormatter{
			FullTimestamp: true,
		},
		Out: out,
	}

	logger.SetLevel(logLevel)
	// logger.SetLevel(log.InfoLevel)
	contextLogger := logger.WithFields(log.Fields{
		"logFrom": name,
	})

	return contextLogger
}
