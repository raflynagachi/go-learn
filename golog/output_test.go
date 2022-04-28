package golog

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutputFile(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)
	logger.Warn("This is Warning Log")
	logger.Trace("This is Trace log")
}
