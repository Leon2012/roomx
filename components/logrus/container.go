package logrus

import (
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
	"github.com/sirupsen/logrus"
)

type Option func(c *Container)

type Container struct {
	config     *Config
	name       string
	coreLogger *elog.Component
	lrLogger     *logrus.Logger
}

func DefaultContainer() *Container {
	return &Container{
		config:     DefaultConfig(),
		coreLogger: elog.EgoLogger.With(elog.FieldComponent(PackageName)),
	}
}

func Load(key string) *Container {
	c := DefaultContainer()
	if err := econf.UnmarshalKey(key, &c.config); err != nil {
		c.coreLogger.Panic("parse config error", elog.FieldErr(err), elog.FieldKey(key))
		return c
	}
	c.coreLogger = c.coreLogger.With(elog.FieldComponentName(key))
	c.name = key
	return c
}

func (c *Container) Build(options ...Option) *Component {
	if options == nil {
		options = make([]Option, 0)
	}
	for _, option := range options {
		option(c)
	}
	component := newComponent(c.name, c.config, c.coreLogger, c.lrLogger)
	return component
}

func WithLogger(logger *logrus.Logger) Option {
	return func(c *Container) {
		c.lrLogger = logger
	}
}
