package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var LogrusLogger *logrus.Logger

func init() {
	LogrusLogger = logrus.New()
	LogrusLogger.SetFormatter(&logrus.JSONFormatter{})
	LogrusLogger.SetOutput(os.Stdout)
	LogrusLogger.SetLevel(logrus.InfoLevel)
}
