package producer

import (
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
	"roomx/libs/queue"
)

const PackageName = "component.queue.producer"

type Component = queue.NsqProducer

type Container struct {
	config *queue.NsqOption
	name   string
	logger *elog.Component
}

func DefaultContainer() *Container {
	return &Container{
		config: DefaultConfig(),
		logger: elog.EgoLogger.With(elog.FieldComponent(PackageName)),
	}
}

func DefaultConfig() *queue.NsqOption {
	return &queue.NsqOption{
		Addr:                "127.0.0.1:4150",
		LookupdAddrs:        nil,
		DialTimeout:         1000,
		ReadTimeout:         60000,
		WriteTimeout:        1000,
		MaxInFlight:         1,
		QueueMsgSize:        10,
		LookupdPollInterval: 0,
		Concurrency:         0,
	}
}

func Load(key string) *Container {
	c := DefaultContainer()
	if err := econf.UnmarshalKey(key, &c.config); err != nil {
		c.logger.Panic("parse config error", elog.FieldErr(err), elog.FieldKey(key))
		return c
	}

	c.logger = c.logger.With(elog.FieldComponentName(key))
	c.name = key
	return c
}

func (c *Container) Build() *Component {
	var (
		producer *queue.NsqProducer
		err      error
	)
	if producer, err = queue.NewNsqProducer(c.config); err != nil {
		c.logger.Panic("new nsq producer error", elog.FieldErr(err), elog.FieldKey("build"))
	}
	return producer
}
