package logging

import "github.com/sirupsen/logrus"

func NewLogrusLogger() *logrus.Logger {
	// TODO: inject level configuration
	return logrus.New()
}
