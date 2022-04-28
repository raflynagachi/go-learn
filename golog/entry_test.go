package golog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"username": "nagachi",
		"gender":   "male",
	})
	entry.Info("Hello entry")
}
