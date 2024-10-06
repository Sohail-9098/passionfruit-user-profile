package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	Info  = "Info"
	Warn  = "Warn"
	Error = "Error"
	Fatal = "Fatal"
	Panic = "Panic"
)

func SetupLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(file)
}

func Log(level, msg string) {
	logLine := logrus.WithFields(logrus.Fields{
		"timestamp": time.Now(),
	})
	switch level {
	case Info:
		logLine.Info(msg)
	case Warn:
		logLine.Warn(msg)
	case Error:
		logLine.Error(msg)
	case Fatal:
		logLine.Fatal(msg)
	case Panic:
		logLine.Panic(msg)
	}
}
