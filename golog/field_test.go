package golog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogWithField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "nagachi").Info("Hello logger")

	logger.WithField("username", "nagachi").
		WithField("name", "Rafly Rigan Nagachi").
		Info("Hello logger")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "nagachi",
		"name":     "Rafly Rigan Nagachi",
		"city":     "OKU Timur",
	}).Info("Hello logger")
}
