package service

import (
	"github.com/sirupsen/logrus"
)

// sLog used to log the errors and info
// all the logs are logged into the console in top level, instead of logging errors multiple times in multiple layers
func logg() logrus.FieldLogger {
	return logrus.WithField("pkg", "service")
}
