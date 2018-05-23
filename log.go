package goctapusmongo

import (
	Log "github.com/sirupsen/logrus"
)

func InitLogger(logLevel string) {
	SetLogLevel(logLevel)
	customFormatter := new(Log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	Log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}

func SetLogLevel(level string) {
	switch level {
	case "debug":
		Log.SetLevel(Log.DebugLevel)
	case "info":
		Log.SetLevel(Log.InfoLevel)
	case "warning":
		Log.SetLevel(Log.WarnLevel)
	case "error":
		Log.SetLevel(Log.ErrorLevel)
	case "fatal":
		Log.SetLevel(Log.FatalLevel)
	case "panic":
		Log.SetLevel(Log.PanicLevel)
	default:
		Log.SetLevel(Log.WarnLevel)
	}
}
