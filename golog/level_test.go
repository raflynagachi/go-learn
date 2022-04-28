package golog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is Info")
	logger.Warn("This is warn")
	logger.Error("This is error")

	// logger.Fatal("This is fatal")
	// logger.Panic("This is panic")
}

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	// log level Trace ke bawah akan di-print
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is Info")
	logger.Warn("This is warn")
	logger.Error("This is error")

	// logger.Fatal("This is fatal")
	// logger.Panic("This is panic")
}
