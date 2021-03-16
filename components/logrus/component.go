package logrus

import (
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"github.com/sirupsen/logrus"
)

const PackageName = "component.logrus"

type Component struct {
	coreLogger *elog.Component
	name       string
	config     *Config
	logger     *Logger
}

func newComponent(name string, config *Config, coreLogger *elog.Component, lrLogger *logrus.Logger) *Component {
	var (
		logger *Logger
		err    error
	)
	if lrLogger == nil {
		logger, err = newWithConfig(config)
	} else {
		logger = newWithLogger(config, lrLogger)
	}
	if err != nil {
		coreLogger.Panic("init logger failed")
	}
	return &Component{
		coreLogger: coreLogger,
		name:       name,
		config:     config,
		logger:     logger,
	}
}

func (c *Component) Debug(format string, a ...interface{}) {
	c.logger.Debug(fmt.Sprintf(format, a...))
}

func (c *Component) Info(format string, a ...interface{}) {
	c.logger.baseLogger.Info(fmt.Sprintf(format, a...))
}

func (c *Component) Error(format string, a ...interface{}) {
	c.logger.baseLogger.Error(fmt.Sprintf(format, a...))
}

func (c *Component) Fatal(format string, a ...interface{}) {
	c.logger.baseLogger.Fatal(fmt.Sprintf(format, a...))
}

func (c *Component) Warn(format string, a ...interface{}) {
	c.logger.baseLogger.Warn(fmt.Sprintf(format, a...))
}

func (c *Component) Trace(format string, a ...interface{}) {
	c.logger.baseLogger.Trace(fmt.Sprintf(format, a...))
}

func (c *Component) GetLogger() *Logger {
	return c.logger
}
