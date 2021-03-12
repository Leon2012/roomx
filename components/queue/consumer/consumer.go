package consumer

import (
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
	"roomx/libs/queue"
)

const PackageName = "component.queue.consumer"

type Component = queue.NsqConsumer

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
		Topic:               "",
		Channel:             "",
		LookupdAddrs:        []string{"127.0.0.1:4161"},
		DialTimeout:         1000,
		ReadTimeout:         60000,
		WriteTimeout:        1000,
		MaxInFlight:         1,
		QueueMsgSize:        10,
		LookupdPollInterval: 1,
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

func (c *Container) Build(h queue.HandlerGenerator) *Component {
	var (
		consumer *queue.NsqConsumer
		err      error
	)
	if c.config.Topic == "" || c.config.Channel == "" || h == nil {
		c.logger.Panic("error params")
	}
	if consumer, err = queue.NewNsqConsumer(c.config.Topic, c.config.Channel, h, c.config); err != nil {
		c.logger.Panic("new nsq consumer error", elog.FieldErr(err), elog.FieldKey("build"))
	}
	return consumer
}
