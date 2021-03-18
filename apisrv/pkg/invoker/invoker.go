package invoker

import (
	"roomx/components/logrus"
)

var (
	Logger *logrus.Component
)

func Init() error {
	Logger = logrus.Load("logger.logrus").Build()
	return nil
}
