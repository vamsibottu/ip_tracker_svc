package clients

import (
	"github.com/sirupsen/logrus"
)

// Log used to log the errors and info
// all the logs are logged into the console in top level, instead of logging errors multiple times in multiple layers
// top level of this micro service is api layer
func logg() logrus.FieldLogger {
	return logrus.WithField("pkg", "clients")
}
